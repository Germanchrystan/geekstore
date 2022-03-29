package domain

type Session struct {
	ID        string `json:"_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
