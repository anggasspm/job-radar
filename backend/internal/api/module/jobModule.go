package module

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/api/rest/handlers"
	"github.com/anggasspm/job-radar/backend/internal/repository"
	"github.com/anggasspm/job-radar/backend/internal/routes"
	"github.com/anggasspm/job-radar/backend/internal/service"
)

func SetupJobModule(rh *rest.RestHandler) {
	jobRepo := repository.NewJobRepository(rh.DB)
	jobService := service.NewJobService(jobRepo)
	jobHandler := handlers.NewJobHandler(jobService)
	routes.SetupJobRoutes(rh, jobHandler)
}
