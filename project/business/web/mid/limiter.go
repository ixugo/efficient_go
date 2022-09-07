package mid

import (
	"github.com/gin-gonic/gin"
)

// ConnLimiter ...
type ConnLimiter struct {
	bucker chan struct{}
}

// NewConnLimiter ...
func NewConnLimiter(n int) *ConnLimiter {
	return &ConnLimiter{
		bucker: make(chan struct{}, n),
	}
}

// Get ..
func (c *ConnLimiter) Get() bool {
	select {
	case c.bucker <- struct{}{}:
		return true
	default:
		return false
	}
}

// Put ..
func (c *ConnLimiter) Put() {
	select {
	case <-c.bucker:
	default:
	}
}

// Limiter 限流中间件
func Limiter(max int) gin.HandlerFunc {
	l := NewConnLimiter(max)
	return func(c *gin.Context) {
		if !l.Get() {
			// 限流
			c.AbortWithStatusJSON(503, gin.H{"msg": "服务繁忙"})
			return
		}
		c.Next()
		l.Put()
	}
}
