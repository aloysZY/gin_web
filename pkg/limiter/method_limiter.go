package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// MethodLimiter 嵌套通用结构体
type MethodLimiter struct {
	*Limiter
}

// NewMethodLimiter 初始化通用令牌桶
func NewMethodLimiter() LimiterIface {
	l := &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimiter{
		Limiter: l,
	}
}

// Key 对特定的接口进行限流
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

// GetBucket 根据 map key 取令牌
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// AddBuckets 根据 map key 添加令牌
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	/*创建令牌桶的方法：
	// 创建指定填充速率和容量大小的令牌桶
	func NewBucket(fillInterval time.Duration, capacity int64) *Bucket
	// 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
	func NewBucketWithQuantum(fillInterval time.Duration, capacity, quantum int64) *Bucket
	// 创建填充速度为指定速率和容量大小的令牌桶
	// NewBucketWithRate(0.1, 200) 表示每秒填充20个令牌
	func NewBucketWithRate(rate float64, capacity int64) *Bucket*/
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}
	return l
}
