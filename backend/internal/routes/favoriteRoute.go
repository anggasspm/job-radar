package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/middleware"
)

func SetupFavoriteRoutes(rh *rest.RestHandler, h *handlers.FavoriteHandler) {
	favoriteRouter := rh.App.Group("/favorite")

	protected := favoriteRouter.Group("/")
	protected.Use(middleware.Authorize(&rh.Auth))

	{
		protected.GET("/", h.GetFavoritesByUser)
	}
}
