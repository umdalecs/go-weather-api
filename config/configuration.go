package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	RedisAddr    string
	RedisPassw   string
	RedisDB      int
	ServerPort   string
	ApiKey       string
	RequestLimit int
}

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		RedisAddr:    LoadEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassw:   LoadEnv("REDIS_PASSWORD", "redis"),
		RedisDB:      LoadIntEnv("REDIS_DB", "0"),
		ServerPort:   fmt.Sprintf(":%s", LoadEnv("SERVER_PORT", "8080")),
		ApiKey:       LoadEnv("WEATHER_VISUALCROSSING_APIKEY", ""),
		RequestLimit: LoadIntEnv("REQUEST_LIMIT", "10"),
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
