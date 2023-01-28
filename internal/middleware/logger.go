package middleware

import (
	"time"

	"github.com/aloysZy/gin_web/global"
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
		/*		// 读取 body 内容,数据直接读取到缓存，数据太大会内存消耗严重,上传文件的时候内内占用大,等待时间长
				bodyByte, _ := ioutil.ReadAll(c.Request.Body)
				// 将读取的内容重新赋值，不然上面读取后之后的路由不能读取
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))*/
		// 如果为如果太大了就不写入日志,后面分别判断 body 大小，然后如果小于一定值记录，不然记录空？

		c.Next()

		cost := time.Since(start)
		// 获取登录的用户
		_userID, _ := c.Get(global.UserId)
		userID, _ := _userID.(uint64)
		logger.Lg.Info(
			path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			// zap.String("body", string(bodyByte)),
			zap.String("ip", c.ClientIP()),
			zap.Uint64("user", userID),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
