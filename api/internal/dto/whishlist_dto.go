package dto

type Whishlist_Dto struct {
	ID       string                   `json:"_id"`
	Products []ProductDisplayCard_Dto `json:"products"`
}
