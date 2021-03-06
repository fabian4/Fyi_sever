package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fabian4/Fyi_sever/push/config"
	"reflect"
)

type HttpPushClient struct {
	endpoint   string
	appId      string
	token      string
	authClient *AuthClient
	client     *HTTPClient
}

// NewClient creates a instance of the huawei cloud common client
// It's contained in huawei cloud app and provides service through huawei cloud app
func NewHttpClient(c *config.Config) (*HttpPushClient, error) {
	if c.AppId == "" {
		return nil, errors.New("appId can't be empty")
	}

	if c.PushUrl == "" {
		return nil, errors.New("pushUrl can't be empty")
	}

	client, err := NewHTTPClient()
	if err != nil {
		return nil, errors.New("failed to get http client")
	}

	authClient, err := NewAuthClient(c)
	if err != nil {
		return nil, err
	}

	token, err := authClient.GetAuthToken(context.Background())
	if err != nil {
		return nil, errors.New("refresh token fail")
	}

	return &HttpPushClient{
		endpoint:   c.PushUrl,
		appId:      c.AppId,
		token:      token,
		authClient: authClient,
		client:     client,
	}, nil
}

func (c *HttpPushClient) refreshToken() error {
	if c.authClient == nil {
		return errors.New("can't refresh token because getting auth client fail")
	}

	token, err := c.authClient.GetAuthToken(context.Background())
	if err != nil {
		return errors.New("refresh token fail")
	}

	c.token = token
	return nil
}

func (c *HttpPushClient) resetHTTPHeader(request *PushRequest) *PushRequest {
	request.Header = []HTTPOption{
		SetHeader("Content-Type", "application/json;charset=utf-8"),
		SetHeader("Authorization", "Bearer "+c.token),
	}
	return request
}
func (c *HttpPushClient) executeApiOperation(ctx context.Context, request *PushRequest, responsePointer interface{}) error {
	err := c.sendHttpRequest(ctx, request, responsePointer)
	if err != nil {
		return err
	}

	// if need to retry for token timeout or other reasons
	retry, err := c.isNeedRetry(responsePointer)
	if err != nil {
		return err
	}

	if retry {
		c.resetHTTPHeader(request)
		err = c.sendHttpRequest(ctx, request, responsePointer)
		return err
	}
	return err
}

func (c *HttpPushClient) sendHttpRequest(ctx context.Context, request *PushRequest, responsePointer interface{}) error {
	resp, err := c.client.DoHttpRequest(ctx, request)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(resp.Body, responsePointer); err != nil {
		return err
	}
	return nil
}

// if token is timeout or error or other reason, need to refresh token and send again
func (c *HttpPushClient) isNeedRetry(responsePointer interface{}) (bool, error) {
	tokenError, err := isTokenError(responsePointer)
	if err != nil {
		return false, err
	}

	if !tokenError {
		return false, nil
	}

	err = c.refreshToken()
	if err != nil {
		return false, err
	}
	return true, nil
}

// if token is timeout or error, refresh token and send again
func isTokenError(responsePointer interface{}) (bool, error) {
	//the responsePointer must be point of struct
	val, _, ok := checkParamStructPtr(responsePointer)
	if !ok {
		return false, errors.New("the parameter should be pointer of the struct")
	}

	code := val.Elem().FieldByName("Code").String()
	if code == config.TokenTimeoutErr || code == config.TokenFailedErr {
		return true, nil
	}
	return false, nil
}

func checkParamStructPtr(structPtr interface{}) (reflect.Value, reflect.Type, bool) {
	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Ptr {
		fmt.Println("The Parameter should be Pointer of Struct!")
		return reflect.Value{}, nil, false
	}

	t := reflect.Indirect(val).Type()
	if t.Kind() != reflect.Struct {
		fmt.Println("The Parameter should be Pointer of Struct!")
		return reflect.Value{}, nil, false
	}
	return val, t, true
}
