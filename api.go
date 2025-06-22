package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ApiServer struct {
	addr string
	rdb  *redis.Client
}

func NewApiServer(addr string, rdb *redis.Client) *ApiServer {
	return &ApiServer{
		addr: addr,
		rdb:  rdb,
	}
}

var (
	ctx = context.Background()
)

func (s *ApiServer) Run() error {
	e := gin.Default()

	rate, _ := limiter.NewRateFromFormatted(fmt.Sprintf("%d-M", Envs.RequestLimit))
	store := memory.NewStore()
	middleware := ginlimiter.NewMiddleware(limiter.New(store, rate))

	e.Use(middleware)

	e.GET("/:location", func(c *gin.Context) {
		location := c.Param("location")

		val, err := s.rdb.Get(ctx, location).Result()
		if err != nil {
			val, err = RetrieveData(location)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
				return
			}

			_, err = s.rdb.Set(ctx, location, val, 12*time.Hour).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing data"})
				return
			}
		}

		c.Header("Content-Type", "application/json")
		c.Status(200)
		c.Writer.Write([]byte(val))
	})

	return e.Run(":8080")
}
