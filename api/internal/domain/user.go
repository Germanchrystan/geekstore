package domain

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
	IsAdmin   bool   `json:"is_admin"`
	IsBanned  bool   `json:"is_banned"`
}

// Addresses      []Address    `json:"addresses"`
// CreditCardData []CreditCard `json:"creditCardData"`
