package domain

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"category_name"`
	Description string `json:"category_description"`
}
