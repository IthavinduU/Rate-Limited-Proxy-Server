package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	target = "https://httpbin.org"
	proxy  *httputil.ReverseProxy
)

func init() {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("❌ Invalid target URL: %v", err)
	}

	proxy = httputil.NewSingleHostReverseProxy(targetURL)

	// Modify response after proxy forwards
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Printf("✅ Forwarded request to: %s", resp.Request.URL.String())
		return nil
	}

	// Custom error handler
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		log.Printf("❌ Proxy error: %v", err)
		http.Error(w, "Proxy error: "+err.Error(), http.StatusBadGateway)
	}
}

// HandleProxy proxies the HTTP request to the target
func HandleProxy(w http.ResponseWriter, r *http.Request) {
	proxy.ServeHTTP(w, r)
}
