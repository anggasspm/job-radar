package service

import (
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/repository"
)

type JobService struct {
	Repo repository.JobRepository
	// repo
}

func NewJobService(Repo repository.JobRepository) *JobService {
	return &JobService{
		Repo: Repo,
	}
}

func (s *JobService) GetJobs() ([]*domain.Job, error) {
	jobs, err := s.Repo.FindJobs()

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (s *JobService) GetJob(id uint) (*domain.Job, error) {
	job, err := s.Repo.FindJob(id)
	if err != nil {
		return nil, err
	}

	return job, nil

}

func (s *JobService) SearchJobs(keyword string) ([]*domain.Job, error) {
	if keyword == "" {
		return s.Repo.FindJobs()
	}

	jobs, err := s.Repo.SearchJobs(keyword)

	if err != nil {
		return nil, err
	}

	return jobs, nil

}
