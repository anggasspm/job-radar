package repository

import (
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"gorm.io/gorm"
)

type FavoriteRepository interface {
	FindFavsByUser(userId uint) ([]*domain.FavoriteJob, error)
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}

func (f *favoriteRepository) FindFavsByUser(userId uint) ([]*domain.FavoriteJob, error) {
	var favorites []*domain.FavoriteJob

	err := f.db.Table("favorites f").
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
