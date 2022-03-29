package domain

type AddressesUsers struct {
	ID        string `json:"_id"`
	UserID    string `json:"user_id"`
	AddressID string `json:"address_id"`
}
