package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
)

func SetupJobRoutes(rh *rest.RestHandler, h *handlers.JobHandler) {
	rh.App.GET("/jobs", h.GetAllJobs)
}
