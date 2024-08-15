package models

type User struct {
	CommonModel
	Investors []Investor `json:"investors"`
	Locations []Location `json:"locations"`

	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
