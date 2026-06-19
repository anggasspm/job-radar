package rest

import "github.com/gin-gonic/gin"

func ErrorMessage(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

func SuccessMessage(c *gin.Context, msg string, data interface{}) {
	
}
