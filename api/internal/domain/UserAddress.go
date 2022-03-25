package domain

type UserAddress struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	AddressID string `json:"address_id"`
}
