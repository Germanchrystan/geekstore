package domain

type Stock struct {
	Id        int    `json:"id"`
	ProductID int    `sjson:"product_id"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	Quantity  int    `json:"quantity"`
}
