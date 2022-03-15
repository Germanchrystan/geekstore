package domain

type Address struct {
	Street       string `json:"street"`
	StreetNumber string `json:"streetNumber"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Zipcode      string `json:"zipcode"`
}
