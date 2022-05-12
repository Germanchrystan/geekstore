package dto

type Order_Dto struct {
	UserId   string  `json:"user_id"`
	StockId  string  `json:"stock_id"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}
