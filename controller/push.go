package controller

import (
	"github.com/fabian4/Fyi_sever/service"
	"github.com/gin-gonic/gin"
)

type PushApi struct {
}

var pushService = service.PushService{}

func (pushApi *PushApi) Push(c *gin.Context) {
	token := c.Param("token")
	tag := c.Param("tag")
	msg := c.DefaultQuery("msg", "msg")
	detail := c.DefaultQuery("detail", "detail")
	pushService.SinglePush(token, msg, detail, tag)
}
