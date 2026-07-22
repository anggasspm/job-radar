package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/anggasspm/job-radar/backend/internal/domain"
	"github.com/anggasspm/job-radar/backend/internal/dto"
	"github.com/anggasspm/job-radar/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// Register godoc
// @Summary User register
// @Description User register
// @Tags users
// @Produce json
// @Success 200 {object} dto.UserSignupResponse
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	// req is a type of dto user signup
	var req dto.UserSignup

	// bind req json to req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.svc.SignUp(req)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrEmailAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": "email already registered. Try logging in"})
		default:
			log.Printf("signup error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Set cookies for access token
	c.SetCookie("access_token", resp.AccessToken, 15*60, "/", "", false, true)

	// Set cookies for refresh token
	c.SetCookie("refresh_token", resp.RefreshToken, 7*24*60*60, "/", "", false, true)

	// response if success
	c.JSON(http.StatusCreated, gin.H{
		"message": "Register successful",
		// "data": &dto.UserSignupResponse{
		// 	User:         resp.User,
		// 	AccessToken:  resp.AccessToken,
		// 	RefreshToken: resp.RefreshToken,
		// },
	})
}

// Login godoc
// @Summary User login
// @Description User login
// @Tags users
// @Produce json
// @Success 200 {object} dto.UserSigninResponse
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLogin

	// bind json
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.svc.Login(req)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserNotFound):
			c.JSON(http.StatusConflict, gin.H{"error": "email not found. Try register"})
		default:
			log.Printf("signup error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Set cookies for access token
	c.SetCookie("access_token", resp.AccessToken, 15*60, "/", "", false, true)

	// Set cookies for refresh token
	c.SetCookie("refresh_token", resp.RefreshToken, 7*24*60*60, "/", "", false, true)

	// response if success
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		// "data": &dto.UserSigninResponse{
		// 	User:         resp.User,
		// 	AccessToken:  resp.AccessToken,
		// 	RefreshToken: resp.RefreshToken,
		// },
	})

}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Refresh token not found",
		})
		return
	}

	resp, err := h.svc.RefreshToken(refreshToken)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserNotFound):
			c.JSON(http.StatusConflict, gin.H{"error": "Token not found. Try register"})
		default:
			log.Printf("signup error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Set cookies for access token
	c.SetCookie("access_token", resp.AccessToken, 15*60, "/", "", false, true)

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed",
		// "data": &dto.RefreshTokenResponse{
		// 	AccessToken: resp.AccessToken,
		// },
	})
}
