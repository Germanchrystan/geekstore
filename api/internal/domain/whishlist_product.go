package domain

type WhishlistProduct struct {
	ID          string `json:"_id"`
	WhishlistID string `json:"whishlist_id"`
	ProductID   string `json:"product_id"`
}
