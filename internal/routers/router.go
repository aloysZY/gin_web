// Package routers 总路由
package routers

import (
	"net/http"

	"github.com/aloysZy/gin_web/internal/middleware"
	v1 "github.com/aloysZy/gin_web/internal/routers/app/v1"
	"github.com/aloysZy/gin_web/pkg/logger"
	"github.com/gin-gonic/gin"

	_ "github.com/aloysZy/gin_web/docs" // 千万不要忘了导入把你上一步生成的docs
	gs "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles" //这个项目被下面的替换了
	"github.com/swaggo/files"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // default就初始化两个中间件了
	// r := gin.New() // 使用自己的中间件
	// 添加日志后设置中间件,添加翻译器中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Translations())
	// swagger 路由
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 初始化，以后 api 版本变更，直接更换初始化的方法就行了
	tag := v1.NewTag()
	// 创建路由组
	apiV1 := r.Group("/api/v1")
	{
		// 设计路由的时候，使用不同的方法进行不同的操作
		apiV1.POST("/tags", tag.Create)            // 创建
		apiV1.GET("/tags", tag.Get)                // 获取
		apiV1.DELETE("/tags/:id", tag.Delete)      // 删除
		apiV1.PUT("/tags/:id", tag.Update)         // 全量更新
		apiV1.PATCH("/tags/:id/state", tag.Update) // 更新部分

		apiV1.POST("/articles")            // 创建
		apiV1.GET("/articles")             // 获取
		apiV1.DELETE("/articles/:id")      // 删除
		apiV1.PUT("/articles/:id")         // 全量更新
		apiV1.PATCH("/articles/:id/state") // 更新部分

	}
	// 没有路由匹配
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}