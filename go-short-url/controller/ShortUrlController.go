package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-short-url/model/param"
	"go-short-url/service"
	"go-short-url/tools"
)

type ShortUrlController struct {
}

func (short *ShortUrlController) Router(router *gin.Engine) {
	userGroup := router.Group("/shortUrl")
	userGroup.POST("/addOrUpdate", short.addOrUpdate)
}

func (short *ShortUrlController) addOrUpdate(context *gin.Context) {
	var param param.UpdateShortUrlParam
	if err := context.BindJSON(&param); err != nil {
		fmt.Println("绑定参数异常")
		return
	}
	tools.Log.Error("ShortUrlController->addOrUpdate,入参=", param)
	service.ShortUrlService{}.AddOrUpdate(&param)
}
