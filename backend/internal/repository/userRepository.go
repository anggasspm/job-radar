package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"gorm.io/gorm"
)

// todo, need to be using same return pattern
type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
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
	if err := r.db.Create(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// still 2, but other wrapped inside fmt.error
			return domain.User{}, fmt.Errorf("create user %s: %w", u.Email, domain.ErrEmailAlreadyExists)
		}
		return domain.User{}, fmt.Errorf("create user %s: %w", u.Email, err)
	}
	return u, nil
}

func (r *userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}
