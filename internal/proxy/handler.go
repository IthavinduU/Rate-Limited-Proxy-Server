package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func HandleProxy(w http.ResponseWriter, r *http.Request) {
	// Change this to any target service you want to forward to
	targetURL := "https://httpbin.org"

	url, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Bad target URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Modify the request before forwarding
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Printf(" Forwarded request to: %s%s", url, r.URL.Path)
		return nil
	}

	// Optional: Custom error handler
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		log.Printf(" Proxy error: %v", err)
		http.Error(w, "Proxy error", http.StatusBadGateway)
	}

	proxy.ServeHTTP(w, r)
}
