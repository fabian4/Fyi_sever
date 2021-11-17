package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fabian4/Fyi_sever/push/model"
	"net/http"
)

func (c *HttpPushClient) SendMessage(ctx context.Context, msgRequest *model.MessageRequest) (*model.MessageResponse, error) {
	result := &model.MessageResponse{}

	request, err := c.getSendMsgRequest(msgRequest)
	if err != nil {
		return nil, err
	}

	err = c.executeApiOperation(ctx, request, result)
	if err != nil {
		return result, err
	}

	b, _ := json.Marshal(msgRequest)
	fmt.Println(string(b))

	return result, err
}

func (c *HttpPushClient) getSendMsgRequest(msgRequest *model.MessageRequest) (*PushRequest, error) {
	body, err := json.Marshal(msgRequest)
	if err != nil {
		return nil, err
	}

	request := &PushRequest{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(c.endpoint, c.appId),
		Body:   body,
		Header: []HTTPOption{
			SetHeader("Content-Type", "application/json;charset=utf-8"),
			SetHeader("Authorization", "Bearer "+c.token),
		},
	}
	return request, nil
}
