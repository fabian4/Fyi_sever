package service

import (
	"context"
	"fmt"
	"github.com/fabian4/Fyi_sever/push"
	"github.com/fabian4/Fyi_sever/push/config"
	"github.com/fabian4/Fyi_sever/push/model"
)

type PushService struct {
}

func (p *PushService) DoPush() {
	msgRequest := model.NewNotificationMsgRequest()
	msgRequest.Message.Token = []string{"IQAAAACy0nMgAACA0d9tunlH0XbgAO4DsBdgFnIbGdy-MHf8Q_IDGcyPi9AJuMG-T1k4sSDQFQTcwdPyvksuzNbRrEk4TRpzJfj9se91eL0VRRkqLQ"}
	msgRequest.Message.Android = model.GetDefaultAndroid()
	msgRequest.Message.Android.Notification = model.GetDefaultAndroidNotification()

	client := push.Client

	resp, err := client.SendMessage(context.Background(), msgRequest)
	if err != nil {
		fmt.Printf("Failed to send message! Error is %s\n", err.Error())
		return
	}

	if resp.Code != config.Success {
		fmt.Printf("Failed to send message! Response is %+v\n", resp)
		return
	}
}
