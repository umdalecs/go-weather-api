package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/umdalecs/weather-api/config"
)

func InitRedisDb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Envs.RedisAddr,
		Password: config.Envs.RedisPassw,
		DB:       config.Envs.RedisDB,
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
