package routers

import (
	"github.com/fabian4/Fyi_sever/controller"
	"github.com/gin-gonic/gin"
)

var baseApi = controller.BaseApi{}

func InitRouter(r *gin.Engine) {
	r.GET("/", baseApi.Push)
}
