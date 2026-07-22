package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/middleware"
)

func SetupJobRoutes(rh *rest.RestHandler, h *handlers.JobHandler) {
	jobRouter := rh.App.Group("/jobs")

	usingMiddleware := jobRouter.Group("/")
	usingMiddleware.Use(middleware.RateLimiter(rh.Limiter, "jobs"))

	usingMiddleware.GET("/", h.GetAllJobs)
	usingMiddleware.GET("/search", h.SearchJobs)
	usingMiddleware.GET("/:id", h.GetJobById)
}
