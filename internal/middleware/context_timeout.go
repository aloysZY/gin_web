package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// ContextTimeout 想做一个全局的超时控制
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
