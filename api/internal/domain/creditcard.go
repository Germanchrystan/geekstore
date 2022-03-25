package domain

type CreditCard struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	HashedData      string `json:"hashedData"`
	LastCodeNumbers int    `json:"last_code_number"`
}
