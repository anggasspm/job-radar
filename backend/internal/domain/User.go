package domain

type User struct {
	ID        uint
	Email     string
	Password_hash  string
	Name      string
	AvatarUrl *string
	Tier      string 
}
