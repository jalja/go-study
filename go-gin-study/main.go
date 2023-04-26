package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-study/controller"
)

func main() {
	// 1.创建路由
	engine := gin.Default()
	registerUserRouter(engine)
	engine.Run()
}
func registerUserRouter(router *gin.Engine) {
	new(controller.UserController).Router(router)
}
