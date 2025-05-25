package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/IthavinduU/go-rate-limit-proxy/internal/metrics"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		duration := time.Since(start)

		// Log to console
		log.Printf("📥 %s %s from %s [%d] in %v",
			r.Method, r.URL.Path, r.RemoteAddr, rec.status, duration)

		// Prometheus Metrics
		metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
		metrics.ResponseStatus.WithLabelValues(http.StatusText(rec.status)).Inc()
	})
}
