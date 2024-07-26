package models

type Address struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	UserID     uint `json:"userId"`
	InvestorID uint `json:"investorId"`
	LocationID uint `json:"locationId"`

	City    string `json:"city"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	ZipCode string `json:"zipCode"`
}
