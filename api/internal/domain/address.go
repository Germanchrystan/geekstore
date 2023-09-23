package domain

type Address struct {
	Id           int     `json:"id"`
	Street       string  `json:"street"`
	StreetNumber float32 `json:"street_number"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	Zipcode      string  `json:"zipcode"`
}
