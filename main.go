package main

import (
	"fmt"
	"github.com/fabian4/Fyi_sever/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

func main() {
	r := gin.Default()
	routers.InitPushRouter(r)
	cfg, _ := ini.Load("config.ini")
	port := ":" + cfg.Section("sever").Key("port").String()
	if err := r.Run(port); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
