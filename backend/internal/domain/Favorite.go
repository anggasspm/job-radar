package domain

import "time"

type FavoriteJob struct {
	ID        uint      `gorm:"column:id"`
	UserID    uint      `gorm:"column:user_id"`
	JobID     int64     `gorm:"column:job_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (FavoriteJob) TableName() string {
	return "favorites"
}
