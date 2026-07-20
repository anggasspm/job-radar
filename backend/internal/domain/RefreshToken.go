package domain

import "time"

type RefreshToken struct {
	ID         uint      `gorm:"column:id"`
	UserID     uint      `gorm:"column:user_id"`
	Token_hash string    `gorm:"column:token_hash"`
	ExpiresAt  time.Time `gorm:"column:expires_at"`
	Revoked    bool      `gorm:"column:revoked;default:false"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}
