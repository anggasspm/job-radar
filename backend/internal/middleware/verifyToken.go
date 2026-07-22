package middleware

import (
	"net/http"

	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func AuthorizeAccessToken(a helper.Auth) gin.HandlerFunc {

	return func(c *gin.Context) {

		accessToken, err := c.Cookie("access_token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Access token not found",
			})
			return
		}

		user, err := a.VerifyToken(accessToken)

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

