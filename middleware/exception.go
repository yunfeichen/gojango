package middleware

import (
	"fmt"
	"gjango/utils/logger"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

/*
说明：异常捕获中间件，用于返回客户端响应
*/
func Exception(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			// 判断异常是否为正常的请求抛出
			resp, ok := err.(RspInfo)
			if ok {
				// 正常，则返回数据
				ctx.JSON(200, resp)
				ctx.Abort()
				return
			}

			// 不正常，则记录异常日志，并且返回服务器异常
			logger.Error(fmt.Sprintf("未知错误：%s\n详细信息：%v", err, string(debug.Stack())))

			resp = RspInfo{
				Code:    500,
				Message: "500错误",
				Data:    map[string]interface{}{},
			}
			ctx.JSON(200, resp)
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
