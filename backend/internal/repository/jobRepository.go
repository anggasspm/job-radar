package repository

import (
	"errors"
	"fmt"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"gorm.io/gorm"
)

type JobRepository interface {
	FindJobs() ([]*domain.Job, error)
	FindJob(id uint) (*domain.Job, error)
	SearchJobs(keyword string) ([]*domain.Job, error)
}

type jobRepository struct {
	db *gorm.DB
}

// Inject repo to service in router
func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{
		db: db,
	}
}

func (r *jobRepository) FindJobs() ([]*domain.Job, error) {
	var jobs []*domain.Job

	err := r.db.Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *jobRepository) FindJob(id uint) (*domain.Job, error) {
	var job *domain.Job

	err := r.db.First(&job, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find job: %w", domain.ErrJobNotFound)
		}
		return nil, err
	}

	return job, nil
}

func (r *jobRepository) SearchJobs(keyword string) ([]*domain.Job, error) {
	var jobs []*domain.Job

	query := "%" + keyword + "%"
	err := r.db.Where("title ILIKE ? OR company ILIKE ? OR location ILIKE ? OR category ILIKE ?", query, query, query, query).
		Order("created_at DESC").
		Limit(50).
		Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil
}
