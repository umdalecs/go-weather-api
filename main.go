package main

import (
	"log"
)

func main() {
	rdb := InitRedisDb()

	s := NewApiServer(Envs.ServerPort, rdb)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
