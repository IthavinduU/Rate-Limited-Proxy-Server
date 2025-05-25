package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// TotalRequests tracks the number of incoming requests labeled by HTTP method and path.
var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "proxy_total_requests",
		Help: "Total number of requests received",
	},
	[]string{"method", "path"},
)

// ResponseStatus tracks the number of responses by status code.
var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "proxy_response_status",
		Help: "HTTP response status codes",
	},
	[]string{"status"},
)

// RateLimitExceeded tracks how many requests have been blocked by rate limiting, labeled by IP.
var RateLimitExceeded = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "proxy_rate_limit_exceeded",
		Help: "Requests that were rate-limited",
	},
	[]string{"ip"},
)

// Init initializes and registers Prometheus metrics.
func Init() {
	prometheus.MustRegister(TotalRequests, ResponseStatus, RateLimitExceeded)
}
