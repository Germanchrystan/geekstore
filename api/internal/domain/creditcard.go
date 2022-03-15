package domain

type CreditCard struct {
	ID         string `json:"id"`
	User       User   `json:"user"`
	HashedData string `json:"hashedData"`
}
