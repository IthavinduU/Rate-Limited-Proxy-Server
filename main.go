package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/IthavinduU/go-rate-limit-proxy/internal/metrics"
	"github.com/IthavinduU/go-rate-limit-proxy/internal/middleware"
	"github.com/IthavinduU/go-rate-limit-proxy/internal/proxy"
)

func main() {
	// Initialize Prometheus metrics
	metrics.Init()

	// Define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", proxy.HandleProxy)
	mux.Handle("/metrics", promhttp.Handler())

	// Wrap with middleware: logging -> rate limit -> final handler
	handler := middleware.LoggingMiddleware(
		middleware.RateLimitMiddleware(mux),
	)

	// Start server
	log.Println(" Proxy server started on http://localhost:8080")
	log.Println(" Prometheus metrics available at http://localhost:8080/metrics")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf(" Failed to start server: %v", err)
	}
}
