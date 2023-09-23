package domain

type User struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Email          string `json:"email"`
	IsActive       bool   `json:"is_active"`
	HashedPassword string `json:"hashed_password"`
	IsAdmin        bool   `json:"is_admin"`
	IsBanned       bool   `json:"is_banned"`
}
