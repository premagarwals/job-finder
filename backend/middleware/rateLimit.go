package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	Requests map[string]int
	Limit    int
	Duration time.Duration
	mu       sync.Mutex
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		Requests: make(map[string]int),
		Limit:    limit,
		Duration: duration,
	}
	go rl.cleanup()
	return rl
}

func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.Duration)
		rl.mu.Lock()
		for k := range rl.Requests {
			delete(rl.Requests, k)
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) LimitRequests(c *gin.Context) {
	ip := c.ClientIP()
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if _, exists := rl.Requests[ip]; !exists {
		rl.Requests[ip] = 0
	}

	if rl.Requests[ip] >= rl.Limit {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
		c.Abort()
		return
	}

	rl.Requests[ip]++
	c.Next()
}
