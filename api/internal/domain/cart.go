package domain

type Cart struct {
	Id        int     `json:"id"`
	UserID    int     `json:"user_id"`
	AddressID int     `json:"address_id"`
	State     string  `json:"state"` //Use enum
	Total     float32 `json:"total"`
}
