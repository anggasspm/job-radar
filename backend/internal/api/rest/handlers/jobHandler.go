package handlers

import (
	"github.com/anggasspm/job-radar/backend/internal/api/rest"
	"github.com/anggasspm/job-radar/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	svc service.JobService
	// job service
}

func NewJobHandler(svc service.JobService) *JobHandler {
	return &JobHandler{
		svc: svc,
	}
}

func (h *JobHandler) GetAllJobs(c *gin.Context) {
	jobs, err := h.svc.GetJobs()

	if err != nil {
		rest.ErrorMessage(c, 404, err)
		return
	}

	c.JSON(200, jobs)
}
