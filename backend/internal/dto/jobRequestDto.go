package dto

import (
	"time"

	"github.com/anggasspm/job-radar/backend/internal/domain"
)

type JobResponse struct {
	ID         int64   `json:"id"`
	SourceID   int     `json:"sourceId"`
	ExternalID *string `json:"externalId"`

	Title       string  `json:"title"`
	Company     string  `json:"company"`
	Location    string  `json:"location"`
	Category    *string `json:"category"`
	Description *string `json:"description"`

	SalaryMin int64  `json:"salaryMin"`
	SalaryMax int64  `json:"salaryMax"`
	Currency  string `json:"currency"`

	MinExp    int16   `json:"minExp"`
	MaxExp    int16   `json:"maxExp"`
	Education *string `json:"education"`

	RawURL string `json:"rawUrl"`

	PostedDate *time.Time `json:"postedDate"`
	ScrapedAt  time.Time  `json:"scrapedAt"`

	IsActive bool `json:"isActive"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToJobResponse(j *domain.Job) JobResponse {
	return JobResponse{
		ID:          j.ID,
		SourceID:    j.SourceID,
		ExternalID:  j.ExternalID,
		Title:       j.Title,
		Company:     j.Company,
		Location:    j.Location,
		Category:    j.Category,
		Description: j.Description,
		SalaryMin:   j.SalaryMin,
		SalaryMax:   j.SalaryMax,
		Currency:    j.Currency,
		MinExp:      j.MinExp,
		MaxExp:      j.MaxExp,
		Education:   j.Education,
		RawURL:      j.RawURL,
		PostedDate:  j.PostedDate,
		ScrapedAt:   j.ScrapedAt,
		IsActive:    j.IsActive,
		CreatedAt:   j.CreatedAt,
		UpdatedAt:   j.UpdatedAt,
	}
}

func ToJobDto(jobs []*domain.Job) []JobResponse {
	result := make([]JobResponse, len(jobs))

	for i, job := range jobs {
		result[i] = ToJobResponse(job)
	}

	return result
}
