package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
)

func SetupUserRoutes(rh *rest.RestHandler, h *handlers.UserHandler) {
	userRouter := rh.App.Group("/auth")
	userRouter.POST("/register", h.Register)
	userRouter.POST("/login", h.Login)
}
