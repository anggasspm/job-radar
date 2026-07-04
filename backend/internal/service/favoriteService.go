package service

import (
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/repository"
)

type FavoriteService struct {
	Repo repository.FavoriteRepository
}

func NewFavService(Repo repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{
		Repo: Repo,
	}
}

func (s *FavoriteService) GetFavsByUser(userId uint) ([]*domain.FavoriteJob, error) {
	favs, err := s.Repo.FindFavsByUser(userId)

	if err != nil {
		return nil, err
	}

	return favs, nil

}
