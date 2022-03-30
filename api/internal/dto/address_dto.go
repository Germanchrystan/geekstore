package dto

type InputAddress_Dto struct {
	SessionID    string  `json:"session_id"`
	Street       string  `json:"street"`
	StreetNumber float32 `json:"street_number"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	Zipcode      string  `json:"zipcode"`
}

type RemoveAddress_Dto struct {
	SessionID string `json:"session_id"`
	AddressID string `json:"address_id"`
}
