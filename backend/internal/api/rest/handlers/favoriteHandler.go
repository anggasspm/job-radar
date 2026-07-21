package handlers

import (
	"net/http"
	"strconv"

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

	favoritesByUser, err := h.svc.GetFavsByUser(c.Request.Context(), uint(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, favoritesByUser)
}

func (h *FavoriteHandler) AddToFavorites(c *gin.Context) {
	userId := c.GetUint("userID")

	idParam := c.Param("id")

	jobId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || jobId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	fav := &dto.FavoriteRequest{
		UserID: userId,
		JobID:  jobId,
	}

	favorite, err := h.svc.AddToFavs(c.Request.Context(), fav)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, favorite)

}

func (h *FavoriteHandler) DeleteFromFavorites(c *gin.Context) {
	userId := c.GetUint("userID")

	idParam := c.Param("id")

	jobId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || jobId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	fav := &dto.FavoriteRequest{
		UserID: userId,
		JobID:  jobId,
	}

	err = h.svc.DeleteFromFavs(c.Request.Context(), fav)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}
