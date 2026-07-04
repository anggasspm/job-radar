package dto

import (
	"time"

	"github.com/anggasspm/job-radar/backend/internal/domain"
)

type FavoriteResponse struct {
	ID        uint      `json:"id"`
	JobID     int64     `json:"job_id"`
	Title     string    `json:"title"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	SalaryMin int64     `json:"salary_min"`
	SalaryMax int64     `json:"salary_max"`
	CreatedAt time.Time `json:"created_at"`
}

func ToFavResponse(f *domain.FavoriteJob) FavoriteResponse {
	return FavoriteResponse{
		ID:        f.ID,
		JobID:     f.JobID,
		Title:     f.Title,
		Company:   f.Company,
		Location:  f.Location,
		SalaryMin: f.SalaryMin,
		SalaryMax: f.SalaryMax,
	}
}

func ToFavDto(favs []*domain.FavoriteJob) []FavoriteResponse {
	result := make([]FavoriteResponse, len(favs))

	for i, fav := range favs {
		result[i] = ToFavResponse(fav)
	}

	return result
}
