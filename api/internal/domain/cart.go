package domain

type Cart struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	State  string `json:"state"` //Use state

}
