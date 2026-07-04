package domain

import "time"

type FavoriteJob struct {
	ID        uint      `gorm:"column:favorite_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	JobID     int64     `gorm:"column:job_id"`
	Title     string    `gorm:"column:title"`
	Company   string    `gorm:"column:company"`
	Location  string    `gorm:"column:location"`
	SalaryMin int64     `gorm:"column:salary_min"`
	SalaryMax int64     `gorm:"column:salary_max"`
}
