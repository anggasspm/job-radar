package middleware

import (
	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func Authorize(a *helper.Auth) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "missing authorization header",
			})
			return
		}

		user, err := a.VerifyToken(authHeader)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("user", user)
		c.Set("userID", user.ID)

		c.Next()
	}
}
