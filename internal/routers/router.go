package routers

import (
	"net/http"

	"github.com/aloysZy/gin_web/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// r := gin.Default() // default就初始化两个中间件了
	r := gin.New() // 使用自己的中间件
	// 添加日志后设置中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 创建路由组
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/test", func(c *gin.Context) {
			c.JSONP(http.StatusOK, gin.H{"message": "test statusOK"})
		})
	}
	// 没有路由匹配
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
