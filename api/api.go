package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type ApiServer struct {
	Addr string
	Rdb  *redis.Client
}

func NewApiServer(addr string, rdb *redis.Client) *ApiServer {
	return &ApiServer{
		Addr: addr,
		Rdb:  rdb,
	}
}

var (
	ctx = context.Background()
)

func (s *ApiServer) Run() error {
	r := http.NewServeMux()

	r.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		WriteJson(w, http.StatusNotFound, map[string]any{"error": "Not found"})
	})

	r.HandleFunc("GET /{location}", func(w http.ResponseWriter, r *http.Request) {
		location := r.PathValue("location")

		val, err := s.Rdb.Get(ctx, location).Result()
		if err != nil {
			val, err = RetrieveData(location)
			if err != nil {
				WriteJson(w, http.StatusInternalServerError, map[string]any{"error": "Error retrieving data"})
				return
			}

			_, err = s.Rdb.Set(ctx, location, val, 12*time.Hour).Result()
			if err != nil {
				WriteJson(w, http.StatusInternalServerError, map[string]any{"error": "Error storing data"})
				return
			}
		}

		WriteStringJson(w, http.StatusOK, val)

	})

	middlewareChain := MiddlewareChain(
		RequestLogger,
		RateLimiter,
	)

	httpServer := &http.Server{
		Addr:    s.Addr,
		Handler: middlewareChain(r),
	}

	log.Printf("Server running in port %s", s.Addr)

	return httpServer.ListenAndServe()
}
