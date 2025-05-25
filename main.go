package main

import (
	"log"
	"net/http"

	"github.com/IthavinduU/go-rate-limit-proxy/internal/proxy"
)

func main() {
	mux := http.NewServeMux()

	// Proxy handler
	mux.HandleFunc("/", proxy.HandleProxy)

	log.Println("🔁 Proxy server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf(" Failed to start server: %v", err)
	}
}
