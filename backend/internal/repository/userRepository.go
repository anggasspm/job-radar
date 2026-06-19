package repository

import "gorm.io/gorm"

type UserRepository interface {
	// all the repo assign here
}

type userRepository struct {
	db *gorm.DB
}

// Inject repo to service in router
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// query using gorm