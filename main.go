package main

import (
	"fmt"
	"github.com/fabian4/Fyi_sever/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
