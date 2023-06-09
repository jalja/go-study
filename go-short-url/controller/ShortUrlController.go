package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-short-url/model/param"
	"go-short-url/model/response"
	"go-short-url/service"
	"go-short-url/tools"
	"net/http"
)

type ShortUrlController struct {
}

func (short *ShortUrlController) Router(router *gin.Engine) {
	userGroup := router.Group("/shortUrl")
	userGroup.POST("/getByPage", short.getByPage)
	userGroup.POST("/addOrUpdate", short.addOrUpdate)
	userGroup.GET("/turnUrl/:url", short.turnUrl)
}

func (short *ShortUrlController) getByPage(context *gin.Context) {
	var param param.ShortUrlParam
	if err := context.BindJSON(&param); err != nil {
		context.JSON(http.StatusOK, response.Error("解析入参异常"))
		return
	}
	tools.Log.Error("getByPage,入参=", param)
	context.JSON(http.StatusOK, response.Success(service.ShortUrlService{}.PageList(&param)))
}

func (short *ShortUrlController) addOrUpdate(context *gin.Context) {
	var param param.UpdateShortUrlParam
	if err := context.BindJSON(&param); err != nil {
		fmt.Println("绑定参数异常")
		return
	}
	tools.Log.Error("ShortUrlController->addOrUpdate,入参=", param)
	context.JSON(http.StatusOK, response.Success(service.ShortUrlService{}.AddOrUpdate(&param)))
}

/** 根据短地址查询实际地址并重定向 **/
func (short *ShortUrlController) turnUrl(context *gin.Context) {
	nativeUrl, err := service.ShortUrlService{}.Redirect(context.Param("url"))
	if err != nil {
		context.JSON(http.StatusOK, response.Error(err.Error()))
		return
	}
	context.Redirect(http.StatusMovedPermanently, nativeUrl)
}
