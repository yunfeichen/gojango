package middleware

import (
	"fmt"
	"gjango/utils"
	"gjango/utils/logger"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

/*
说明：访问日志中间件
*/

func AccessLog(ctx *gin.Context) {

	// 请求路由
	requestUri := ctx.Request.RequestURI

	// 开始时间
	requestStartTime := time.Now()

	if strings.HasPrefix(requestUri, "/static") {
		//查看普通会员用户的下载大小，进行控制
		//errinfo := ginrsp.ResponseInfo{
		//	Code:    450,
		//	Message: "static AccessLog 鉴权失败",
		//}
		//ctx.JSON(200, errinfo)
		//ctx.Abort()
		//return
	}

	// 处理请求
	ctx.Next()
	// 结束时间
	requestEndTime := time.Now()
	// 处理耗时
	requestExecTime := requestEndTime.Sub(requestStartTime)
	// 请求方式
	requestMethod := ctx.Request.Method

	// 状态码
	requestCode := ctx.Writer.Status()
	// 请求 IP
	requestIP := ctx.ClientIP()

	responseStatus := ctx.Writer.Status()
	////响应 header
	//responseHeader := ctx.Writer.Header()
	//响应体大小
	responseBodySize := ctx.Writer.Size()

	if strings.HasPrefix(requestUri, "/static") {
		//存储普通会员用户返回响应的size
		logger.Debug(fmt.Sprintf("%s\t%d\t%d", requestUri, responseStatus, responseBodySize))
	}

	// 判断请求方式，OPTIONS 使用 DEBUG 输出
	if requestMethod == "OPTIONS" {
		logger.Debug(fmt.Sprintf("%s\t%s\t%d\t%s\t%s", requestMethod, requestUri, requestCode, requestExecTime.String(), requestIP))
	} else {
		logger.Info(fmt.Sprintf("%s\t%s\t%d\t%s\t%s", requestMethod, requestUri, requestCode, requestExecTime.String(), requestIP))
	}
}

func FileAccessLog(ctx *gin.Context) {
	// 处理请求
	ctx.Next()
	fmt.Println("进入FileAccessLog")
	requestUri := ctx.Request.RequestURI
	responseStatus := ctx.Writer.Status()
	////响应 header
	//responseHeader := ctx.Writer.Header()
	//响应体大小
	responseBodySize := ctx.Writer.Size()
	logger.Debug(fmt.Sprintf("%s\t%d\t%d", requestUri, responseStatus, responseBodySize))

}

func JwtParseInfo(ctx *gin.Context) {

	// 请求方式
	requestMethod := ctx.Request.Method
	// 请求路由
	requestUri := ctx.Request.RequestURI

	fmt.Println("JwtParseInfo method:" + requestMethod + ",URL:" + requestUri)

	token_s := ctx.Request.Header["Token"][0]
	fmt.Println("JwtParseInfo token_s:", token_s)

	token, err := jwt.Parse(token_s, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetConfig().Jwt.Key), nil
	})

	if err != nil {
		fmt.Println("parse token 失败！")
		fmt.Println("err:" + err.Error())
	}
	if !token.Valid {
		fmt.Println("token无效")

	} else {
		fmt.Println("token有效")
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("转为claim失败")
		} else {
			sub := claim["sub"].(string)
			fmt.Println("sub is ", sub)
		}

	}

	// 处理请求
	ctx.Next()

}
