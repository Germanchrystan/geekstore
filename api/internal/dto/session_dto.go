package dto

import "github.com/Germanchrystan/GeekStore/api/internal/domain"

type Session_Dto struct {
	Session     domain.Session          `json:"session"`
	User        domain.User             `json:"user"`
	Adresses    []domain.Address        `json:"addresses"`
	CreditCards []DisplayCreditCard_Dto `json:"credit_cards"`
}
