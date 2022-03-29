package domain

type Address struct {
	ID           string  `json:"_id"`
	Street       string  `json:"street"`
	StreetNumber float32 `json:"street_number"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	Zipcode      string  `json:"zipcode"`
}
