package helper

import (
	"errors"
	"time"

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


