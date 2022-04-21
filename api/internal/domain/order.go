package domain

type Order struct {
	ID       string  `json:"_id"`
	StockID  string  `json:"stock_id"`
	CartID   string  `json:"cart_id"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"order_price"`
}
