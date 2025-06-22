package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedisDb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Envs.RedisAddr,
		Password: Envs.RedisPass,
		DB:       Envs.RedisDB,
	})

	if err := testConnection(rdb); err != nil {
		log.Fatal(err)
	}

	return rdb
}

func testConnection(rdb *redis.Client) error {
	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Print("Connected to Redis.")

	return nil
}
