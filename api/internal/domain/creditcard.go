package domain

type CreditCard struct {
	Id              int    `json:"id"`
	UserID          int    `json:"user_id"`
	HashedData      string `json:"hashed_data"`
	LastCodeNumbers int    `json:"last_code_numbers"`
}
