package service

import (
	"context"

	"github.com/anggasspm/job-radar/backend/cache"
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/dto"
	"github.com/anggasspm/job-radar/backend/internal/repository"
)

type FavoriteService struct {
	Repo  repository.FavoriteRepository
	Cache cache.FavoriteCache
}

func NewFavService(Repo repository.FavoriteRepository, Cache cache.FavoriteCache) *FavoriteService {
	return &FavoriteService{
		Repo:  Repo,
		Cache: Cache,
	}
}

func (s *FavoriteService) GetFavsByUser(ctx context.Context, userId uint) ([]*dto.FavoriteDetail, error) {
	// get data from cache
	favs, err := s.Cache.Get(ctx, userId)
	if err == nil {
		return favs, nil
	}

	// if data not in cache -> set the data
	favs, err = s.Repo.FindFavsByUser(userId)
	if err != nil {
		return nil, err
	}
	_ = s.Cache.Set(ctx, userId, favs)

	return favs, nil

}

func (s *FavoriteService) AddToFavs(ctx context.Context, req *dto.FavoriteRequest) (*dto.FavoriteResponse, error) {
	fav, err := s.Repo.AddToFavs(&domain.FavoriteJob{
		JobID:  req.JobID,
		UserID: req.UserID,
	})

	if err != nil {
		return nil, err
	}

	_ = s.Cache.Delete(ctx, req.UserID)

	return &dto.FavoriteResponse{
		ID:        fav.ID,
		UserID:    fav.UserID,
		JobID:     fav.JobID,
		CreatedAt: fav.CreatedAt,
	}, nil
}

func (s *FavoriteService) DeleteFromFavs(ctx context.Context, req *dto.FavoriteRequest) error {
	err := s.Repo.DeleteFromFavs(&domain.FavoriteJob{
		JobID:  req.JobID,
		UserID: req.UserID,
	})

	if err != nil {
		return err
	}

	_ = s.Cache.Delete(ctx, req.UserID)

	return nil
}
