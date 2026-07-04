package handlers

import (
	"net/http"

	"github.com/anggasspm/job-radar/backend/internal/dto"
	"github.com/anggasspm/job-radar/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	svc *service.FavoriteService
}

func NewFavoriteHandler(svc *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		svc: svc,
	}
}

func (h *FavoriteHandler) GetFavoritesByUser(c *gin.Context) {
	userId := c.GetUint("userID")

	favoritesByUser, err := h.svc.GetFavsByUser(uint(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToFavDto(favoritesByUser))
}

func (h *FavoriteHandler) AddToFavorites(c *gin.Context) {

}

func (h *FavoriteHandler) DeleteFromFavorites(c *gin.Context) {

}
