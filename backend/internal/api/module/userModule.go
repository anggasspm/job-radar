package module

import (
	"github.com/anggasspm/job-radar/backend/config"
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/anggasspm/job-radar/backend/internal/repository"
	"github.com/anggasspm/job-radar/backend/internal/routes"
	"github.com/anggasspm/job-radar/backend/internal/service"
)

func SetupUserModule(rh *rest.RestHandler, cfg *config.AppConfig) {
	auth := helper.SetupAuth(cfg.AppSecret)
	userRepo := repository.NewUserRepository(rh.DB)
	userService := service.NewUserService(userRepo, auth)
	userHandler := handlers.NewUserHandler(userService)
	routes.SetupUserRoutes(rh, userHandler)
}
