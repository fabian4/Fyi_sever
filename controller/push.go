package controller

import (
	"github.com/fabian4/Fyi_sever/service"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

var pushService = service.PushService{}

func (baseApi *BaseApi) Push(c *gin.Context) {
	pushService.DoPush()
}
