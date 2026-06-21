package domain

// gorm tag not json tag
type User struct {
	ID            uint
	Email         string `gorm:"uniqueIndex;not null"`
	Password_hash string `gorm:"column:password_hash"`
	Name          string
	AvatarUrl     *string
	Tier          string `gorm:"type:varchar(10);default:free"`
}
