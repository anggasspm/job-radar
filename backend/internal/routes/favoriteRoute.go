package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/middleware"
)

func SetupFavoriteRoutes(rh *rest.RestHandler, h *handlers.FavoriteHandler) {
	favoriteRouter := rh.App.Group("/favorite")

	protected := favoriteRouter.Group("/")
	protected.Use(middleware.AuthorizeAccessToken(&rh.Auth), middleware.RateLimiter(rh.Limiter, "favorite"))

	{
		protected.GET("/", h.GetFavoritesByUser)
		protected.POST("/:id", h.AddToFavorites)
		protected.DELETE("/:id", h.DeleteFromFavorites)
	}
}
