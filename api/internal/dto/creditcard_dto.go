package dto

// This DTO struct is used to store Credit Cards
type InputCreditCard_Dto struct {
	SessionID    string `json:"session_id"`
	Code         string `json:"code"`
	ExpiryDate   string `json:"expiry_date"`
	SecurityCode int    `json:"security_code"`
}

type DisplayCreditCard_DTO struct {
	LastCodeNumbers int `json:"last_code_number"`
}
