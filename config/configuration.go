package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	RedisAddr  string
	RedisPassw string
	RedisDB    int

	ApiKey string

	RequestLimit int
}

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		RedisAddr:  LoadEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassw: LoadEnv("REDIS_PASSWORD", "redis"),
		RedisDB:    LoadIntEnv("REDIS_DB", "0"),

		RequestLimit: LoadIntEnv("REQUEST_LIMIT", "10"),

		ApiKey: LoadEnv("WEATHER_VISUALCROSSING_APIKEY", ""),
	}
}

func LoadEnv(name, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return fallback
}

func LoadIntEnv(name, fallback string) int {
	value, _ := strconv.Atoi(LoadEnv(name, fallback))
	return value
}
