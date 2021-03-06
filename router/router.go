package router

import (
	"gin-demo/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 添加 Get 请求路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.GET("/add/:a/:b", handler.AddResult)
	r.GET("/sub", handler.SubResult)
	r.POST("/mul", handler.MulResult)
	r.POST("/div", handler.DivResult)
	r.POST("/sum", handler.SumResult)
	return r
}
