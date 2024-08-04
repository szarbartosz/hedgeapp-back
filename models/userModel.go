package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Investors []Investor `json:"investors"`
	Locations []Location `json:"locations"`

	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
