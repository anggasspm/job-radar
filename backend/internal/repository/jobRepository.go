package repository

import (
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

func (j *jobRepository) FindJobs() ([]*domain.Job, error) {
	var jobs []*domain.Job

	err := j.db.Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (j *jobRepository) FindJob(id uint) (*domain.Job, error) {
	var job *domain.Job

	err := j.db.First(&job, id).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (j *jobRepository) SearchJobs(keyword string) ([]*domain.Job, error) {
	var jobs []*domain.Job

	query := "%" + keyword + "%"
	err := j.db.Where("title ILIKE ? OR company ILIKE ? OR location ILIKE ? OR category ILIKE ?", query, query, query, query).
		Order("created_at DESC").
		Limit(50).
		Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil
}
