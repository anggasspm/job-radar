package module

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/repository"
	"github.com/anggasspm/job-radar/backend/internal/routes"
	"github.com/anggasspm/job-radar/backend/internal/service"
)

func SetupFavoriteModule(rh *rest.RestHandler) {
	favoriteRepo := repository.NewFavoriteRepository(rh.DB)
	favoriteService := service.NewFavService(favoriteRepo)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)
	routes.SetupFavoriteRoutes(rh, favoriteHandler)

}
