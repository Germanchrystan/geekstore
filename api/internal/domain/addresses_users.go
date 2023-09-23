package domain

type AddressesUsers struct {
	Id        int `json:"id"`
	UserID    int `json:"user_id"`
	AddressID int `json:"address_id"`
}
