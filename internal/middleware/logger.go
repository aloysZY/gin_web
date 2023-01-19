package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/aloysZy/gin_web/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 接收gin框架默认的日志，添加访问日志记录
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		// 读取 body 内容,数据直接读取到缓存，数据太大会内存消耗严重
		bodyByte, _ := ioutil.ReadAll(c.Request.Body)
		// 将读取的内容重新赋值，不然上面读取后之后的路由不能读取
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
		c.Next()

		cost := time.Since(start)
		logger.Lg.Info(
			path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("body", string(bodyByte)),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
