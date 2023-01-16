package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // default就初始化两个中间件了
	// 创建路由组
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/test", func(c *gin.Context) {
			c.JSONP(http.StatusOK, gin.H{"message": "test statusOK"})
		})
	}

	return r
}
