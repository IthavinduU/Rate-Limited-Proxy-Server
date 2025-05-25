package middleware

import (
	"net"
	"net/http"
	"sync"

	"github.com/IthavinduU/go-rate-limit-proxy/internal/metrics"
	"github.com/IthavinduU/go-rate-limit-proxy/internal/ratelimiter"
)

var (
	ipBuckets = make(map[string]*ratelimiter.TokenBucket)
	mutex     sync.Mutex
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)

		bucket := getTokenBucket(ip)
		if !bucket.Allow() {
			metrics.RateLimitExceeded.WithLabelValues(ip).Inc()
			http.Error(w, "429 - Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getTokenBucket(ip string) *ratelimiter.TokenBucket {
	mutex.Lock()
	defer mutex.Unlock()

	bucket, exists := ipBuckets[ip]
	if !exists {
		bucket = ratelimiter.NewTokenBucket(10, 1) // capacity=10, refillRate=1 token/sec
		ipBuckets[ip] = bucket
	}
	return bucket
}

func getIP(r *http.Request) string {
	// Try X-Forwarded-For header first (useful if behind reverse proxy)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return xff
	}

	// Fallback: use remote address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // fallback
	}
	return ip
}
