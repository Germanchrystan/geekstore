package dto

type Login_Dto struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}
