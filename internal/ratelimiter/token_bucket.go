package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket implements a thread-safe token bucket rate limiter.
type TokenBucket struct {
	capacity     int       // Max number of tokens
	tokens       int       // Current number of tokens
	refillRate   int       // Tokens per second
	lastRefilled time.Time // Last refill timestamp
	mutex        sync.Mutex
}

// NewTokenBucket creates a new TokenBucket with given capacity and refill rate.
func NewTokenBucket(capacity int, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

// Allow checks if a request is allowed by the rate limiter.
func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefilled).Seconds()

	// Calculate how many tokens to refill
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

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
