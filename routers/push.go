package routers

import (
	"github.com/fabian4/Fyi_sever/controller"
	"github.com/gin-gonic/gin"
)

var pushApi = controller.PushApi{}

func InitPushRouter(r *gin.Engine) {
	r.GET("/:token/push/:tag", pushApi.Push)
}
