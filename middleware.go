package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type clientData struct {
	requests int
	reset    time.Time
}

var (
	clients = make(map[string]*clientData)
	mu      sync.Mutex
	limit   = Envs.RequestLimit
	window  = 1 * time.Minute
)

type Middleware func(http.Handler) http.Handler

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next
	}
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		mu.Lock()
		data, exists := clients[ip]
		if !exists || time.Now().After(data.reset) {
			clients[ip] = &clientData{requests: 0, reset: time.Now().Add(window)}
		} else {
			data.requests++

			if data.requests >= limit {
				mu.Unlock()
				WriteJson(w, http.StatusTooManyRequests, map[string]string{"error": "Too many requests"})
				return
			}
		}
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
