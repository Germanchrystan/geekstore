package domain

type WhishlistProduct struct {
	Id          int `json:"id"`
	WhishlistID int `json:"whishlist_id"`
	ProductID   int `json:"product_id"`
}
