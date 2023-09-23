package domain

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"product_name"`
	Price       float32 `json:"price"`
	Description string  `json:"product_description"`
	SubgenreId  int     `json:"subgenre_id"`
	CategoryId  int     `json:"category_id"`
}
