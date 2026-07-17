package service

import (
	"fmt"
	"time"

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
// can add more context to error from helper rather than just generic error from helper
func (s *UserService) SignUp(req dto.UserSignup) (*dto.UserSignupResponse, error) {
	passwordHash, err := s.Auth.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user, err := s.Repo.CreateUser(&domain.User{
		Email:         req.Email,
		Password_hash: passwordHash,
		Name:          req.Name,
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

	hashedToken, err := s.Auth.HashToken(accessToken)
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(24 * time.Hour * 7)

	err = s.Repo.SaveRefreshToken(&domain.RefreshToken{
		UserID:     user.ID,
		Token_hash: hashedToken,
		ExpiresAt:  expiresAt,
	})

	return &dto.UserSignupResponse{
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// login
func (s *UserService) Login(req dto.UserLogin) (*dto.UserSigninResponse, error) {
	user, err := s.Repo.FindUserByEmail(req.Email)

	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	err = s.Auth.VerifyPassword(req.Password, user.Password_hash)

	if err != nil {
		return nil, fmt.Errorf("eror on verifying password %w", err)
	}

	accessToken, err := s.Auth.GenerateAccessToken(user.ID, user.Email, user.Tier)
	if err != nil {
		return nil, fmt.Errorf("error on generating access token %w", err)
	}

	refreshToken, err := s.Auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("error on generating refresh token %w", err)
	}

	return &dto.UserSigninResponse{
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
