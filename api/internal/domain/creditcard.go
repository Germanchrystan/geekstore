package domain

type CreditCard struct {
	ID              string `json:"_id"`
	UserID          string `json:"user_id"`
	HashedData      string `json:"hashed_data"`
	LastCodeNumbers int    `json:"last_code_number"`
}
