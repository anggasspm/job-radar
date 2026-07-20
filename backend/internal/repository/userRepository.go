package repository

import (
	"errors"
	"fmt"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"gorm.io/gorm"
)

// todo, need to be using same return pattern
type UserRepository interface {
	CreateUser(u *domain.User) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	FindUserById(id uint) (*domain.User, error)
	SaveRefreshToken(rf *domain.RefreshToken) error
	FindToken(rf *domain.RefreshToken) (*domain.RefreshToken, error)
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
func (r *userRepository) CreateUser(u *domain.User) (*domain.User, error) {
	if err := r.db.Create(u).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// still 2, but other wrapped inside fmt.error
			return &domain.User{}, fmt.Errorf("create user %s: %w", u.Email, domain.ErrEmailAlreadyExists)
		}
		return nil, fmt.Errorf("create user %s: %w", u.Email, err)
	}
	return u, nil
}

// need to check other errors rather than just label it as not foun
func (r *userRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find user %s: %w", email, domain.ErrUserNotFound)
		}
		return nil, fmt.Errorf("find user %s: %w", email, err)
	}
	return &user, nil
}

func (r *userRepository) FindUserById(id uint) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "id=?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find user %s: %w", id, domain.ErrUserNotFound)
		}
		return nil, fmt.Errorf("find user %s: %w", id, err)
	}
	return &user, nil
}

// Save the refreshToken to the database
func (r *userRepository) SaveRefreshToken(rf *domain.RefreshToken) error {
	if err := r.db.Create(rf).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return nil
}

func (r *userRepository) FindToken(rf *domain.RefreshToken) (*domain.RefreshToken, error) {
	var token domain.RefreshToken

	err := r.db.First(&token, "token_hash=?", rf.Token_hash).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find user %s: %w", rf, err)
		}
		return nil, fmt.Errorf("find user %s: %w", rf, err)
	}

	return &token, nil
}
