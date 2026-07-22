package security

import (
	"github.com/anggasspm/job-radar/backend/cache"
	redisrate "github.com/go-redis/redis_rate/v10"
)

type RedisRateLimiter struct {
	*redisrate.Limiter
}

func NewRedisRateLimiter(redisClient *cache.RedisClient) *RedisRateLimiter {
	return &RedisRateLimiter{
        Limiter: redisrate.NewLimiter(redisClient.RedisClient),
	}
}
