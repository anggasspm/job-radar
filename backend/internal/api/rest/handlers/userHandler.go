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
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// Function
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

	// change user to token, after signup return jwt token
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

	// response if success
	c.JSON(http.StatusCreated, resp)
}
