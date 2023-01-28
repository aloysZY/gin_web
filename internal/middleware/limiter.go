package middleware

import (
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// RateLimiter 根据接口限流桶实现
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	/*取出令牌的方法如下：
	// 取token（非阻塞）
	func (tb *Bucket) Take(count int64) time.Duration
	func (tb *Bucket) TakeAvailable(count int64) int64

	// 最多等maxWait时间取token
	func (tb *Bucket) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool)

	// 取token（阻塞）
	func (tb *Bucket) Wait(count int64)
	func (tb *Bucket) WaitMaxDuration(count int64, maxWait time.Duration) bool*/
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
