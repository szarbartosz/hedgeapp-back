package models

type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
