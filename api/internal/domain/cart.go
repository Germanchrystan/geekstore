package domain

type Cart struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	AddressID string  `json:"address_id"`
	State     string  `json:"state"` //Use enum
	Total     float32 `json:"total"`
}
