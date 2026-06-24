package api

import (
	"log"
	"time"

	"github.com/anggasspm/job-radar/backend/config"
	"github.com/anggasspm/job-radar/backend/internal/api/module"
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(rh *rest.RestHandler, cfg *config.AppConfig) {
	module.SetupUserModule(rh, cfg)
	module.SetupJobModule(rh)
}

func StartServer(cfg config.AppConfig) {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://jobsradars.vercel.app", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	app.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRoutes(rh, &cfg)

	log.Printf("Server running on :%s", cfg.ServerPort)
	app.Run(":" + cfg.ServerPort)
}
