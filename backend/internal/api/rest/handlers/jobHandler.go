package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/dto"
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

// GetAllJobs godoc
// @Summary Get all jobs
// @Description Get all available jobs
// @Tags jobs
// @Produce json
// @Success 200 {array} dto.JobResponse
// @Router /jobs [get]
func (h *JobHandler) GetAllJobs(c *gin.Context) {
	jobs, err := h.svc.GetJobs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToJobDto(jobs))
}

// GetJobById godoc
// @Summary Get job by id
// @Description Get job detail by id
// @Tags jobs
// @Produce json
// @Success 200 {object} dto.JobResponse
// @Router /jobs/:id [get]
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

	c.JSON(http.StatusOK, dto.ToJobResponse(job))
}

// SearchJobs godoc
// @Summary SearchJobNl
// @Description Search jobs with matching query/filter
// @Tags jobs
// @Produce json
// @Success 200 {array} dto.JobResponse
// @Param q query string false "Search keyword"
// @Router /jobs/search [get]
func (h *JobHandler) SearchJobs(c *gin.Context) {
	query := c.Query("q")

	jobs, err := h.svc.SearchJobs(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ToJobDto(jobs))
}
