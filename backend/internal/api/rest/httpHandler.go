package rest

import (
	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RestHandler struct {
	App   *gin.Engine
	DB    *gorm.DB
	Auth  helper.Auth
	Redis *redis.Client
}
