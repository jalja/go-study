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
	userGroup.POST("/addOrUpdate", short.addOrUpdate)
	userGroup.GET("/turnUrl/:url", short.turnUrl)
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
	nativeUrl := service.ShortUrlService{}.GetByShortUrl(context.Param("url"))
	if nativeUrl == "" {
		context.JSON(http.StatusOK, response.Error("地址不存在"))
		return
	}
	context.Redirect(http.StatusMovedPermanently, nativeUrl)
}
