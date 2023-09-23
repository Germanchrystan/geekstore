package domain

type Order struct {
	Id       int     `json:"id"`
	StockID  int     `json:"stock_id"`
	CartID   int     `json:"cart_id"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"order_price"`
}
