package middleware

import (
	"github.com/aloysZy/gin_web/global"
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

const (
	LOCALE = "locale"
	ZH     = "zh"
	EN     = "en"
)

// Translations 翻译器
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化英文、中文、中文简体翻译器
		uni := ut.New(en.New(), zh.New(), zh_Hant.New())
		// 请求头获取字段
		locale := c.GetHeader(LOCALE)
		// 初始化翻译器
		trans, _ := uni.GetTranslator(locale)
		// 修改gin框架中的Validator引擎属性，实现自定制
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case ZH:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
			case EN:
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
			default:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
			}
			// 添加到上下文，后面接口校验的时候会进行使用
			c.Set(global.Trans, trans)
		}
		// 继续路由
		c.Next()
	}
}
