package domain

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	IsActive       bool   `json:"is_active"`
	IsAdmin        bool   `json:"is_admin"`
	IsBanned       bool   `json:"is_banned"`
}
