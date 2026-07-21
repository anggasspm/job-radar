package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

// injector

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}
func (a *Auth) HashPassword(p string) (string, error) {
	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		return "", errors.New("Password hash failed")
	}

	return string(hashP), nil
}

func (a *Auth) HashToken(t string) (string, error) {
	hash := sha256.Sum256([]byte(t))
	return hex.EncodeToString(hash[:]), nil
}

func (a Auth) VerifyPassword(pP string, hP string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a *Auth) GenerateAccessToken(id uint, email string, tier string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"tier":    tier,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenStr, err := accessToken.SignedString([]byte(a.Secret))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (a *Auth) GenerateRefreshToken(id uint) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenStr, err := refreshToken.SignedString([]byte(a.Secret))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// create a jwt verifyToken and authorize that is contains gin.context because router only accepts gin handler func
func (a *Auth) VerifyToken(t string) (*domain.User, error) {
	// verify
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return &domain.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.Tier = claims["tier"].(string)
		return &user, nil
	}

	return &domain.User{}, errors.New("token verification failed")
}

func (a *Auth) VerifyRefreshToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.Secret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid refresh token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	return uint(userID), nil
}
