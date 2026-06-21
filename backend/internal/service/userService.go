package service

import (
	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/dto"
	"github.com/anggasspm/job-radar/backend/internal/helper"
	"github.com/anggasspm/job-radar/backend/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func NewUserService(Repo repository.UserRepository, auth helper.Auth) *UserService {
	return &UserService{
		Repo: Repo,
		Auth: auth,
	}
}

// change to dto
func (s *UserService) SignUp(req dto.UserSignup) (*dto.UserSignupResponse, error) {
	passwordHash, err := s.Auth.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:         req.Email,
		Password_hash: passwordHash,
		Name:          req.Name,
		Tier:          "free",
	})

	if err != nil {
		return nil, err
	}

	accessToken, err := s.Auth.GenerateAccessToken(user.ID, user.Email, user.Tier)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.Auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.UserSignupResponse{
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
