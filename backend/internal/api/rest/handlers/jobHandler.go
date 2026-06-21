package handlers

import (
	"net/http"

	"github.com/anggasspm/job-radar/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	svc *service.JobService
	// job service
}

func NewJobHandler(svc *service.JobService) *JobHandler {
	return &JobHandler{
		svc: svc,
	}
}

func (h *JobHandler) GetAllJobs(c *gin.Context) {
	jobs, err := h.svc.GetJobs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobs)
}
