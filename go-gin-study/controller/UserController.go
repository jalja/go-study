package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-study/model/response"
	"net/http"
)

type UserController struct {
}

func (user *UserController) Router(engine *gin.Engine) {
	//engine.Group 路由组 (http://127.0.0.1:8080/user/getUser)
	userGroup := engine.Group("/user")
	userGroup.GET("/getUserById", user.getUserById)
}

// 更具条件查询 UserInfo
func (user *UserController) getUserById(context *gin.Context) {
	id := context.Param("id")
	fmt.Println("====getOrder=====", id)
	context.JSON(http.StatusOK, response.Success(""))
}
