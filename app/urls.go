package app

import (
	"gjango/api"
	"gjango/middleware"
	"log"

	_ "gjango/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 路由初始化
func InitRouter(g *gin.Engine) {
	g.GET("/swagger/*any", func(c *gin.Context) {
		ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
	})

	g.Use(middleware.AccessLog) // 访问中间件
	g.Use(middleware.Cors)      // 跨域中间件
	g.Use(middleware.Exception) // 异常中间件
	g.Use(gin.Recovery())

	var vGroup *gin.RouterGroup
	vGroup = g.Group("api")
	api.InitRouter(vGroup)

	log.Print("路由初始化完成")

	return
}
