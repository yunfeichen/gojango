package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test() string {
	fmt.Println("hello this is a test")
	return "abc"
}

// @Summary 打印测试功能
// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
// @Host 127.0.0.1:8080
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
// @Router / [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetCookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
}
