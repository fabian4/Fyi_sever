package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fabian4/Fyi_sever/push"
	"github.com/fabian4/Fyi_sever/push/config"
	"github.com/fabian4/Fyi_sever/push/model"
)

type PushService struct {
}

func (p *PushService) DoPush() {
	msgRequest, err := getNotifyMsgRequest()
	if err != nil {
		fmt.Printf("Failed to get message request! Error is %s\n", err.Error())
		return
	}

	client := push.GetPushClient()

	resp, err := client.SendMessage(context.Background(), msgRequest)
	if err != nil {
		fmt.Printf("Failed to send message! Error is %s\n", err.Error())
		return
	}

	if resp.Code != config.Success {
		fmt.Printf("Failed to send message! Response is %+v\n", resp)
		return
	}

	fmt.Printf("Succeed to send message! Response is %+v\n", resp)
}

func getNotifyMsgRequest() (*model.MessageRequest, error) {
	msgRequest := model.NewNotificationMsgRequest()
	msgRequest.Message.Token = push.TargetTokenArray
	msgRequest.Message.Android = model.GetDefaultAndroid()
	msgRequest.Message.Android.Notification = model.GetDefaultAndroidNotification()

	b, err := json.Marshal(msgRequest)
	if err != nil {
		fmt.Printf("Failed to marshal the default message! Error is %s\n", err.Error())
		return nil, err
	}

	fmt.Printf("Default message is %s\n", string(b))
	return msgRequest, nil
}
