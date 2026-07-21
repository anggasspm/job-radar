package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Anything that goes here, means that they're from env
type AppConfig struct {
	ServerPort string
	Dsn        string // Connection string
	AppSecret  string
	RedisHost  string
	RedisPort  string
}

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("APP_PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	Dsn := os.Getenv("DSN")

	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	AppSecret := os.Getenv("APP_SECRET")

	if len(AppSecret) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	RedisHost := os.Getenv("REDIS_HOST")

	if len(RedisHost) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	RedisPort := os.Getenv("REDIS_PORT")

	if len(RedisPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: AppSecret, RedisHost: RedisHost, RedisPort: RedisPort}, nil

}
