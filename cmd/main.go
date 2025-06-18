package main

import (
	"log"

	"github.com/umdalecs/weather-api/api"
	"github.com/umdalecs/weather-api/db"
)

func main() {
	rdb := db.InitRedisDb()

	s := api.NewApiServer(":8080", rdb)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
