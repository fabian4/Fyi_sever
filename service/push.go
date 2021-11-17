package service

import (
	"context"
	"fmt"
	"github.com/fabian4/Fyi_sever/push"
	"github.com/fabian4/Fyi_sever/push/config"
)

type PushService struct {
}

func (p *PushService) SinglePush(token string, msg string, detail string, tag string) {
	msgRequest := push.GetSimpleMessage(msg, detail, tag)
	msgRequest.Message.Token = []string{token}

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
