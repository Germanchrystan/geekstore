package domain

type Session struct {
	ID        string `json:"id"`
	User      User   `json:"user"`
	CreatedAt string `json:"createdAt"`
}
