// Package routers 总路由
package routers

import (
	"net/http"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/internal/middleware"
	"github.com/aloysZy/gin_web/internal/routers/api"
	v1 "github.com/aloysZy/gin_web/internal/routers/api/v1"
	"github.com/gin-gonic/gin"

	_ "github.com/aloysZy/gin_web/docs" // 千万不要忘了导入把你上一步生成的docs
	gs "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles" //这个项目被下面的替换了
	"github.com/swaggo/files"
)

// NewRouter 初始化路由
func NewRouter() *gin.Engine {
	// 设置gin模式，要在初始化之前
	gin.SetMode(global.ServerSetting.RunMode)
	// r := gin.Default() // default就初始化两个中间件了
	r := gin.New() // 使用自己的中间件
	// 添加日志后设置中间件,添加翻译器中间件
	if global.ServerSetting.RunMode != "release" {
		r.Use(gin.Logger())
	}
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout.ContextTimeout)) // 统一超时中间件
	r.Use(middleware.RateLimiter(global.AuthMethodLimiters))                          // 令牌桶
	r.Use(middleware.GinLogger())                                                     // 日志中间件
	r.Use(middleware.GinRecovery())                                                   // recovery中间件
	r.Use(middleware.Translations())                                                  // 翻译器
	r.Use(middleware.Tracing())                                                       // 要在s所有路由注册之前使用 路由追踪
	// swagger 路由
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 初始化，以后 api 版本变更，直接更换初始化的方法就行了
	tag := v1.NewTag()
	upload := v1.NewUpload()
	atricle := v1.NewArticle()
	// 鉴权路由
	auth := api.NewAuth()
	r.POST("/signup", auth.SignUp)
	r.POST("/auth", auth.Auth)
	// 创建路由组
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.Auth())
	{
		apiV1.POST("/upload/file", upload.UploadFile)
		apiV1.StaticFS("/static", http.Dir(global.AppSetting.UploadImage.UploadSavePath))
		// 设计路由的时候，使用不同的方法进行不同的操作
		apiV1.POST("/tags", tag.Create) // 创建
		apiV1.GET("/tags", tag.List)    // 获取标签列表
		// apiV1.GET("/tags", tag.Get)           // 获取单个标签
		// apiV1.GET("/tags", tag.GetById)       // 根据标签 ID 查询标签
		apiV1.DELETE("/tags/:id", tag.Delete) // 删除
		apiV1.PUT("/tags/:id", tag.Update)    // 全量更新
		// apiV1.PATCH("/tags/:id/state", tag.Update) // 更新部分；这个就是改变标签是否可用和PUT重复了
		apiV1.POST("/articles", atricle.Create) // 创建
		apiV1.GET("/articles", atricle.List)    // 获取多个文章
		// apiV1.GET("/articles", atricle.Get)           // 获取单个文章
		// apiV1.GET("/articles", atricle.GetById)       // 根据 ID 获取文章
		apiV1.DELETE("/articles/:id", atricle.Delete) // 删除
		apiV1.PUT("/articles/:id", atricle.Update)    // 全量更新
		// apiV1.PATCH("/articles/:id/state") // 更新部分

	}
	// 没有路由匹配
	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "404",
	// 	})
	// })
	return r
}
