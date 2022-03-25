package domain

type Stock struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	Quantity  int    `json:"quantity"`
}
