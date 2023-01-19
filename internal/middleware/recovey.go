package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/aloysZy/gin_web/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinRecovery recover掉项目可能出现的panic, stack 布尔值来记录堆栈信息
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				// 使用stack 来判断是否记录堆栈信息，现在直接都记录
				logger.Lg.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())))
				// if stack {
				// 	logger.Lg.Error("[Recovery from panic]",
				// 		zap.Any("error", err),
				// 		zap.String("request", string(httpRequest)),
				// 		zap.String("stack", string(debug.Stack())),
				// 	)
				// } else {
				// 	logger.Lg.Error("[Recovery from panic]",
				// 		zap.Any("error", err),
				// 		zap.String("request", string(httpRequest)),
				// 	)
				// }
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
