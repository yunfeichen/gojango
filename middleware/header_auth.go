package middleware

import (
	"gjango/utils"
	"gjango/utils/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type RspInfo struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func HeaderAuth(ctx *gin.Context) {

	// 请求方式
	requestMethod := ctx.Request.Method
	// 请求路由
	requestUri := ctx.Request.RequestURI
	logger.Debug("JwtParseInfo method:" + requestMethod + ",URL:" + requestUri)

	{
		sKey := FirstUpper(utils.GetConfig().KeyAuth.KeyName)
		ApiKeyValues, ok := ctx.Request.Header[sKey]
		if ok {
			//keyauth
			if len(ApiKeyValues) > 0 {
				if ApiKeyValues[0] == utils.GetConfig().KeyAuth.KeyValue {
					logger.Debug(utils.GetConfig().KeyAuth.KeyName + " keyauth鉴权OK!")
					ctx.Next()
					return
				} else {
					errinfo := RspInfo{
						Code:    417,
						Message: "keyauth鉴权失败",
					}
					ctx.JSON(200, errinfo)
					ctx.Abort()
					return
				}
			}
		}
	}

	logger.Debug(utils.GetConfig().KeyAuth.KeyName + " is not exist!")
	logger.Debug(ctx.Request.Header)
	{
		tokens, ok := ctx.Request.Header["Token"]
		if ok {
			//keyauth
			if len(tokens) > 0 {
				token_s := ctx.Request.Header["Token"][0]
				logger.Debug("token_s:", token_s)

				token, err := jwt.Parse(token_s, func(token *jwt.Token) (interface{}, error) {
					return []byte(utils.GetConfig().Jwt.Key + "11111"), nil
				})

				if err != nil {
					logger.Debug("parse token 失败！")
					logger.Debug("err:" + err.Error())
				} else {
					if !token.Valid {
						logger.Debug("token无效")
						//errRspInfo := response.ResponseInfo{
						//	ErrCode:    417,
						//	ErrMessage: "token无效",
						//}
						//ctx.JSON(200, errRspInfo)
						//ctx.Abort()
						//return
					} else {
						logger.Debug("token有效")
						//claim, ok := token.Claims.(jwt.MapClaims)
						//if !ok {
						//	fmt.Println("转为claim失败")
						//} else {
						//	sub := claim["sub"].(string)
						//	fmt.Println("sub is ", sub)
						//}
					}
				}
			}
			// 处理请求
			ctx.Next()
			return
		}

	}

	//errRspInfo := ginrsp.ResponseInfo{
	//	ErrCode:    417,
	//	ErrMessage: "未鉴权",
	//}
	//ctx.JSON(200, errRspInfo)
	//ctx.Abort()
	//return
	// 处理请求
	ctx.Next()

}
