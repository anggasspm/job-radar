package service

import (
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/dto"
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

func (s *FavoriteService) GetFavsByUser(userId uint) ([]*dto.FavoriteDetail, error) {
	favs, err := s.Repo.FindFavsByUser(userId)

	if err != nil {
		return nil, err
	}

	return favs, nil

}

func (s *FavoriteService) AddToFavs(req *dto.FavoriteRequest) (*dto.FavoriteResponse, error) {
	fav, err := s.Repo.AddToFavs(&domain.FavoriteJob{
		JobID:  req.JobID,
		UserID: req.UserID,
	})

	if err != nil {
		return nil, err
	}

	return &dto.FavoriteResponse{
		ID:        fav.ID,
		UserID:    fav.UserID,
		JobID:     fav.JobID,
		CreatedAt: fav.CreatedAt,
	}, nil
}
