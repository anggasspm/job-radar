package routes

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/repository"
	"github.com/anggasspm/job-radar/backend/internal/service"
)

func SetupJobRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.JobService{
		Repo: repository.NewJobRepository(rh.DB),
	}

	handler := handlers.NewJobHandler(svc)

	app.GET("/jobs", handler.GetAllJobs)
}

