package api

import (
	"log"

	"github.com/anggasspm/job-radar/backend/config"
	"github.com/anggasspm/job-radar/backend/internal/api/module"
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(rh *rest.RestHandler, cfg *config.AppConfig) {
	module.SetupUserModule(rh, cfg)
	module.SetupJobModule(rh)
}

func StartServer(cfg config.AppConfig) {
	app := gin.Default()

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	log.Println("Database connected!")

	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: helper.SetupAuth(cfg.AppSecret),
	}

	setupRoutes(rh, &cfg)

	log.Println("Server running on :8080")
	app.Run(":8080")
}
