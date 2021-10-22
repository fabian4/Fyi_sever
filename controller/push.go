package controller

import (
	"github.com/fabian4/Fyi_sever/service"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (baseApi *BaseApi) push(c *gin.Context) {
	var pushService = service.PushService{}
	pushService.DoPush()
}
