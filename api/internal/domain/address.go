package domain

type Address struct {
	ID           string  `json:"id"`
	Street       string  `json:"street"`
	StreetNumber float32 `json:"streetNumber"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	Zipcode      string  `json:"zipcode"`
}
