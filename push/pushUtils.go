package push

import (
	"fmt"
	"github.com/fabian4/Fyi_sever/push/config"
	"github.com/fabian4/Fyi_sever/push/core"
	"github.com/go-ini/ini"
	"sync"
)

const (
	token = "IQAAAACy0nMgAACA0d9tunlH0XbgAO4DsBdgFnIbGdy-MHf8Q_IDGcyPi9AJuMG-T1k4sSDQFQTcwdPyvksuzNbRrEk4TRpzJfj9se91eL0VRRkqLQ"
)

var (
	pushClient *core.HttpPushClient
	once       sync.Once
)

var (
	//TargetToken the topic to be subscribed/unsubscribed
	TargetTopic = "topic"

	//TargetCondition the condition of the devices operated
	TargetCondition = "'topic' in topics && ('topic' in topics || 'TopicC' in topics)"

	//TargetToken the token of the device operated
	TargetToken = token

	//TargetTokenArray the collection of the tokens of th devices operated
	TargetTokenArray = []string{TargetToken}
)

func GetPushClient() *core.HttpPushClient {
	once.Do(func() {
		cfg, _ := ini.Load("config.ini")
		client, err := core.NewHttpClient(getPushConf(cfg))
		if err != nil {
			fmt.Printf("Failed to new common client! Error is %s\n", err.Error())
			panic(err)
		}
		pushClient = client
	})

	return pushClient
}

func getPushConf(cfg *ini.File) *config.Config {
	return &config.Config{
		AppId:     cfg.Section("push").Key("appId").String(),
		AppSecret: cfg.Section("push").Key("appSecret").String(),
		AuthUrl:   cfg.Section("push").Key("authUrl").String(),
		PushUrl:   cfg.Section("push").Key("pushUrl").String(),
	}
}
