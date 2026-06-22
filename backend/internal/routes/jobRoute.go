package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
)

func SetupJobRoutes(rh *rest.RestHandler, h *handlers.JobHandler) {
	jobRouter := rh.App.Group("/jobs")
	jobRouter.GET("/", h.GetAllJobs)
	jobRouter.GET("/search", h.SearchJobs)
	jobRouter.GET("/:id", h.GetJobById)
}
