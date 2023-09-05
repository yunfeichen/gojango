package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", Ping)
	r.GET("/cookie", GetCookie)

	log.Print("路由初始化完成")
	return
}
