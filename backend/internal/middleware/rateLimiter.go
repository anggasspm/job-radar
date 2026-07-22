package middleware

import (
	"strconv"
	"time"

	"github.com/anggasspm/job-radar/backend/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

type RateLimitMiddleware struct {
	RedisLimiter *security.RedisRateLimiter
}

func RateLimiter(r *security.RedisRateLimiter, prefix string) gin.HandlerFunc {

	return func(c *gin.Context) {
		var rate int
		var burst int

		if prefix == "login" {
			rate = 5
			burst = 5
		} else {
			rate = 60
			burst = 20
		}

		key := prefix + ":" + c.ClientIP()

		res, _ := r.Limiter.Allow(c, key, redis_rate.Limit{
			Rate:   rate,
			Burst:  burst,
			Period: time.Minute,
		})
		if res.Allowed <= 0 {
			c.Header("Retry-After", strconv.FormatInt(int64(res.RetryAfter.Seconds()), 10))

			// handle rate limit exceeded error
			c.AbortWithStatusJSON(429, gin.H{
				"message": "too many requests",
			})
			return
		}
		c.Next()
	}
}
