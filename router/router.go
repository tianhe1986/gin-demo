package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 添加 Get 请求路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	return r
}