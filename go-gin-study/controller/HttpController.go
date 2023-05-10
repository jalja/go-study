package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-study/tools"
	"io"
	"net/http"
)

type HttpController struct {
}

func (hc *HttpController) Router(router *gin.Engine) {
	//engine.Group 路由组 (http://127.0.0.1:8080/user/getUser)
	userGroup := router.Group("/http")
	userGroup.GET("/getTest", hc.getTest)
	userGroup.GET("/postTest", hc.postTest)
}

func (hc *HttpController) getTest(context *gin.Context) {

	res, err := http.Get("http://localhost:8080/user/getUserById?id=1")
	if err != nil {
		fmt.Println("http 解析异常")
		return
	}
	by, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http 解析异常")
		return
	}
	tools.Log.Error("getTest=", string(by))
}

func (hc *HttpController) postTest(context *gin.Context) {
	client := &http.Client{}
	str := "{\"sku\":\"\",\"styleNo\":\"\",\"supplierName\":\"\",\"purchaseOrderId\":\"\",\"channelNo\":\"\",\"purchaserName\":\"\",\"purchaseBelongName\":\"\",\"payType\":\"\",\"poId\":\"\",\"jstReturnId\":\"\",\"jstReturnName\":\"\",\"returnDate\":\"\",\"returnDateStart\":\"\",\"returnDateEnd\":\"\",\"statusOptName\":\"\",\"statusOptTime\":\"\",\"statusOptTimeStart\":\"\",\"statusOptTimeEnd\":\"\",\"diffType\":2,\"returnProductStatus\":2,\"currentPage\":1,\"pageSize\":50}"
	req, err := http.NewRequest("POST", "https://gateway-qa.doublefs.com/purchase/returnProduct/selectReturnProductPage", bytes.NewReader([]byte(str)))
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Currency", "USD")
	req.Header.Set("Lang", "en")
	req.Header.Set("System-Source", "WEB")
	req.Header.Set("app-id", "0")
	req.Header.Set("Cookie", "device-id=2303090247049991; _ga=GA1.1.1662532365.1678330026; _scid=050dd3b8-3099-43f1-bf1b-5db6e986e3dc; _fbp=fb.1.1678330025895.2139100389; _tt_enable_cookie=1; _ttp=69FSCTfD5OBi2EJU_ZpiFN0l_wE; ab_test_white_id=skipAbWhite; synchronized_cart=1; rec_product_rel_ABParams=; _gcl_au=1.1.1154315245.1678330029; rskxRunCookie=0; rCookie=slcdrhizacfufjhooeho8lf0id6s4; token=KCQeUL7xx32lAO6kpsWCP+El3Pv17vlgxjUFW8XEQQEL5XtBwmTavhAlwWZrjPkB; ab_test_last_token=KCQeUL7xx32lAO6kpsWCP+El3Pv17vlgxjUFW8XEQQEL5XtBwmTavhAlwWZrjPkB; lastRskxRun=1678330453095; tracking_ordernumber=230309025407664175-PAID; _sctr=1|1679241600000; _uetvid=af3973f0be2411ed917b09cd348d9a3c; forterToken=4b7cb081c69f4ea3a7dabaf3d1bbc5d8_1679280771488__UDF43_13ck; _ga_KV25P2G0RF=GS1.1.1679280769.3.1.1679280772.0.0.0; _ga_4R7D7PS633=GS1.1.1679280769.3.1.1679280772.57.0.0; SprAuthorization=splr-aa276d3d9ab949f2a21df535cb8d3675606; Authorization=u-3OQxeEQwx6oWqGg9N5euxtgk2vNR14HzoG005k400IvT; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2210000001695175%22%2C%22first_id%22%3A%22184383d87821467-032ffb4890be67-18525635-1296000-184383d87831b11%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22%24device_id%22%3A%22184383d87821467-032ffb4890be67-18525635-1296000-184383d87831b11%22%7D; SprJSESSIONID=2363D9ACD7308B881046EC21903D5D95")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http 解析异常")
		return
	}
	by, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http 解析异常")
		return
	}
	tools.Log.Error("postTest=", string(by))
}
