package middleware

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 2) // 🔥 strict for testing
		visitors[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("RateLimit HIT 🔥", r.RemoteAddr)

		ip := r.RemoteAddr

		limiter := getLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many requests 🚫", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}