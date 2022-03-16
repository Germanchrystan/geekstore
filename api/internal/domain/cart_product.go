package domain

type CartOrder struct {
	ID      string `json:"id"`
	CartID  string `json:"cart_id"`
	OrderID string `json:"order_id"`
}
