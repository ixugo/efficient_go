package main

import (
	"fmt"

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
	fmt.Println("get ", len(c.bucker))
	select {
	case c.bucker <- struct{}{}:
		return true
	default:
		return false
	}
}

// Put ..
func (c *ConnLimiter) Put() {
	fmt.Println("put ", len(c.bucker))
	select {
	case <-c.bucker:
	default:
	}
}

// Limiter 限流中间件
func Limiter() gin.HandlerFunc {
	l := NewConnLimiter(2)
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
