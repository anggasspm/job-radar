package dto

import (
	"time"
)

type FavoriteResponse struct {
	ID        uint
	JobID     int64
	UserID    uint
	CreatedAt time.Time
}

type FavoriteRequest struct {
	ID        uint  `json:"id"`
	JobID     int64 `json:"job_id"`
	UserID    uint  `json:"user_id"`
	CreatedAt time.Time
}
