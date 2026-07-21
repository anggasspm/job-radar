package cache

import "github.com/redis/go-redis/v9"

type RedisClient struct {
	RedisClient *redis.Client
}

func NewRedis(addr, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisClient{
		RedisClient: client,
	}
}
