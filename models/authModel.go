package models

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
