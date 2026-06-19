package service

import (
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/repository"
)

type JobService struct {
	Repo repository.JobRepository
	// repo
}

func (s *JobService) GetJobs() ([]*domain.Job, error) {
	jobs, err := s.Repo.FindJobs()

	if err != nil {
		return nil, err
	}

	return jobs, nil
}
