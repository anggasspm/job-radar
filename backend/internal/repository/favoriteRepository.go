package repository

import (
	"fmt"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/dto"
	"gorm.io/gorm"
)

type FavoriteRepository interface {
	FindFavsByUser(userId uint) ([]*dto.FavoriteDetail, error)
	AddToFavs(f *domain.FavoriteJob) (*domain.FavoriteJob, error)
	DeleteFromFavs(f *domain.FavoriteJob) error
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}

func (r *favoriteRepository) FindFavsByUser(userId uint) ([]*dto.FavoriteDetail, error) {
	var favorites []*dto.FavoriteDetail

	err := r.db.Table("favorites f").
		Select(`
            f.id AS favorite_id,
            f.created_at,
            j.id AS job_id,
            j.title,
            j.company,
            j.location,
            j.salary_min,
            j.salary_max
        `).
		Joins("JOIN jobs j ON j.id = f.job_id").
		Where("f.user_id = ?", userId).
		Scan(&favorites).Error

	return favorites, err

}

func (r *favoriteRepository) AddToFavs(f *domain.FavoriteJob) (*domain.FavoriteJob, error) {

	if err := r.db.Create(f).Error; err != nil {
		return nil, fmt.Errorf("create favorite job %d: %w", f.JobID, err)
	}

	return f, nil
}

func (r *favoriteRepository) DeleteFromFavs(f *domain.FavoriteJob) error {

	err := r.db.Table("favorites f").Where("user_id = ? AND job_id = ?", f.UserID, f.JobID).Delete(&domain.FavoriteJob{})

	if err != nil {
		return nil
	}

	return nil
}
