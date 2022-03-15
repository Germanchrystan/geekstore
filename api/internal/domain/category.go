package domain

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"category_name"`
	Description string `json:"category_description"`
}
