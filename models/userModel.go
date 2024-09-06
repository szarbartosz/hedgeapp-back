package models

type User struct {
	CommonModel
	Investors []Investor `json:"investors"`
	Locations []Location `json:"locations"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
}
