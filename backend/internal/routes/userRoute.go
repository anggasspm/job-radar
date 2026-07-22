package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/middleware"
)

func SetupUserRoutes(rh *rest.RestHandler, h *handlers.UserHandler) {
	userRouter := rh.App.Group("/auth")

	usingMiddleware := userRouter.Group("/")
	usingMiddleware.Use(middleware.RateLimiter(rh.Limiter, "jobs"))

	usingMiddleware.POST("/register", h.Register)
	usingMiddleware.POST("/login", h.Login)
	usingMiddleware.POST("/refresh", h.RefreshToken)
}
