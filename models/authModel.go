package models

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpCredentials struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
