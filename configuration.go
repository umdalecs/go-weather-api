package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	RedisAddr string
	RedisPass string
	RedisDB   int

	ApiKey string

	RequestLimit int
}

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		RedisAddr: LoadEnv("REDIS_ADDR"),
		RedisPass: LoadEnv("REDIS_PASSWORD"),
		RedisDB:   LoadIntEnv("REDIS_DB"),

		RequestLimit: LoadIntEnv("REQUEST_LIMIT"),

		ApiKey: LoadEnv("WEATHER_VISUALCROSSING_APIKEY"),
	}
}

func LoadEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("missing %s", name)
	}

	return value
}

func LoadIntEnv(name string) int {
	value, err := strconv.Atoi(LoadEnv(name))
	if err != nil {
		log.Fatalf("%s must be an integer", name)
	}

	return value
}
