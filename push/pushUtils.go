package push

import (
	"fmt"
	"github.com/fabian4/Fyi_sever/push/config"
	"github.com/fabian4/Fyi_sever/push/core"
	"github.com/fabian4/Fyi_sever/push/model"
	"github.com/go-ini/ini"
	"sync"
)

var (
	Client *core.HttpPushClient
	once   sync.Once
)

func init() {
	once.Do(func() {
		cfg, _ := ini.Load("config.ini")
		client, err := core.NewHttpClient(getPushConf(cfg))
		if err != nil {
			fmt.Printf("Failed to new common client! Error is %s\n", err.Error())
			panic(err)
		}
		Client = client
	})
}

func GetSimpleMessage(msg string, detail string, tag string) *model.MessageRequest {
	msgRequest := model.NewNotificationMsgRequest()
	msgRequest.Message.Android = model.GetDefaultAndroid()
	msgRequest.Message.Android.Notification = model.GetAndroidNotification(msg, detail, tag)
	return msgRequest
}

func getPushConf(cfg *ini.File) *config.Config {
	return &config.Config{
		AppId:     cfg.Section("push").Key("appId").String(),
		AppSecret: cfg.Section("push").Key("appSecret").String(),
		AuthUrl:   cfg.Section("push").Key("authUrl").String(),
		PushUrl:   cfg.Section("push").Key("pushUrl").String(),
	}
}
