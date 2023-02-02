package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 翻译器不能作为中间件使用，并发的时候，会引发fatal error: concurrent map read and map write
// https://www.cnblogs.com/wanghaostec/p/15037690.html
const (
	ZH = "zh"
	EN = "en"
)

// Translations 翻译器
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化英文、中文、中文简体翻译器
		uni := ut.New(en.New(), zh.New(), zh_Hant.New())
		// 请求头获取字段
		locale := c.GetHeader("locale")
		// 初始化翻译器，根据要翻译的语言
		trans, _ := uni.GetTranslator(locale)
		// 初始化为全局的
		// global.Trans, _ = uni.GetTranslator(locale)
		// 修改gin框架中的Validator引擎属性，实现自定制
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case ZH:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
				break
			case EN:
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 添加到上下文，后面接口校验的时候会进行使用
			c.Set("trans", trans)
		}
		// 继续路由
		c.Next()
	}
}
