package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-study/model/response"
	"go-gin-study/tools"
	"net/http"
)

type UserController struct {
}

func (user *UserController) Router(router *gin.Engine) {
	//engine.Group 路由组 (http://127.0.0.1:8080/user/getUser)
	userGroup := router.Group("/user")
	userGroup.GET("/getUserById", user.getUserById)
}

// 更具条件查询 UserInfo
func (user *UserController) getUserById(context *gin.Context) {
	id := context.Query("id")
	tools.Log.Infof("UserController-getUserById=>id=%s", id)
	context.JSON(http.StatusOK, response.Success(""))
}
