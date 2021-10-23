package controller

import (
	"github.com/fabian4/Fyi_sever/service"
	"github.com/gin-gonic/gin"
)

type PushApi struct {
}

var pushService = service.PushService{}

func (pushApi *PushApi) Push(c *gin.Context) {
	pushService.DoPush()
}
