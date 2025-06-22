package main

import (
	"log"
)

func main() {
	rdb := InitRedisDb()

	s := NewApiServer(":8080", rdb)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
