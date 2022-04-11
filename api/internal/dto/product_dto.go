package dto

import "github.com/Germanchrystan/GeekStore/api/internal/domain"

type ProductDisplayCard_Dto struct {
	ID    string  `json:"_id"`
	Name  string  `json:"product_name"`
	Price float32 `json:"price"`

	Subgenre string `json:"subgenre_id"`
	Category string `json:"category_id"`

	InWhishlist bool `json:"in_whishlist"`
}

type ProductDisplayDetail_Dto struct {
	ID          string  `json:"_id"`
	Name        string  `json:"product_name"`
	Price       float32 `json:"price"`
	Description string  `json:"product_description"`

	Subgenre string `json:"subgenre_id"`
	Category string `json:"category_id"`

	Stocks []domain.Stock `json:"stocks"`

	InWhishlist bool `json:"in_whishlist"`
}
