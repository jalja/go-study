package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-study/controller"
)

func main() {
	// 1.创建路由
	router := gin.Default()
	registerUserRouter(router)
	router.Run()
}
func registerUserRouter(router *gin.Engine) {
	new(controller.UserController).Router(router)
}
