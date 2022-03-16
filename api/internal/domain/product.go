package domain

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"product_name"`
	Price       float32 `json:"price"`
	Description string  `json:"product_description"`
	Subgenre    string  `json:"subgenre_id"`
	Category    string  `json:"category_id"`
}
