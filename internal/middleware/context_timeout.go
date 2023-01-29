package middleware

import (
	"context"
	"time"

	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// ContextTimeout 全局的超时控制
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				// write response and abort the request
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.GatewayTimeout)
				// c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			// cancel to clear resources after finished
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
