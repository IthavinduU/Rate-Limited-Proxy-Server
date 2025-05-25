package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int
	lastRefilled time.Time
	mutex        sync.Mutex
}

func NewTokenBucket(capacity int, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefilled).Seconds()

	refilled := int(elapsed * float64(tb.refillRate))
	if refilled > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+refilled)
		tb.lastRefilled = now
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
