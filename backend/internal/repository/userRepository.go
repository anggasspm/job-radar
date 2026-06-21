package repository

import (
	"errors"
	"log"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
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
func (r *userRepository) CreateUser(u domain.User) (domain.User, error) {
	err := r.db.Create(&u).Error

	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, err
	}

	return u, nil
}

func (r *userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}
