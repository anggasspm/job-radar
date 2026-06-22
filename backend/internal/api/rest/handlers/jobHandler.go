package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/service"
	"github.com/gin-gonic/gin"
)

// add more spesific type error, and in another layer or create error domain
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

// add error handling not found etc
func (h *JobHandler) GetJobById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	job, err := h.svc.GetJob(uint(id))

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrJobNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": "job not found",
			})
		default:
			log.Printf("find job error: %w", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) SearchJobs(c *gin.Context) {
	query := c.Query("q")

	jobs, err := h.svc.SearchJobs(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, jobs)
}
