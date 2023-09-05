package main

import (
	"fmt"
	"gjango/app"
	"gjango/utils"
	"gjango/utils/logger"
	"gjango/utils/shutdown"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] -  \"%s - %s - %s - %d - %s - \"%s\" - %s\"\n",
			param.TimeStamp.Format(time.DateTime),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// 设置运行模式
	gin.SetMode("debug")

	// 初始化路由
	app.InitRouter(router)

	//创建和启动web服务(gin)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", utils.GetConfig().Listen.Port),
		Handler: router,
	}
	go func() {
		logger.Debug("开始启动", srv.Addr)
		err := srv.ListenAndServe()
		// 启动时候如果报错，并且错误不是关闭服务器，则打印日志并退出
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动失败，%s", err.Error())
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close()
}
