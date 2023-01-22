package middleware

import (
	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}
		if token == "" {
			ecode = errcode.NotLogin
		} else {
			// ParseToken 解析的时候是可以解析出来，自定义的 token 字段的
			claims, err := app.ParseToken(token)
			// 将这个字段上下文传递,用当前用户的时候就使用这个 key 获取
			c.Set(global.UserId, claims.UserId)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				// 判断token 是否超时
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
