package push

import (
	"fmt"
	"github.com/fabian4/Fyi_sever/push/config"
	"github.com/fabian4/Fyi_sever/push/core"
	"sync"
)

const (
	// private data ,it's import,please don't let it out
	appId     = "xxxxxx"
	appSecret = "xxxxxx"
	token     = "xxxxxx"

	// below is public address
	// get token address
	authUrl = "https://login.cloud.huawei.com/oauth2/v2/token"
	// send push msg address
	pushUrl = "https://api.push.hicloud.com"
)

var conf = &config.Config{
	AppId:     appId,
	AppSecret: appSecret,
	AuthUrl:   authUrl,
	PushUrl:   pushUrl,
}

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
		client, err := core.NewHttpClient(conf)
		if err != nil {
			fmt.Printf("Failed to new common client! Error is %s\n", err.Error())
			panic(err)
		}
		pushClient = client
	})

	return pushClient
}

func GetPushConf() *config.Config {
	return conf
}
