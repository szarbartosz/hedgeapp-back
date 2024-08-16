package models

type OwnerType string

const (
	InvestorType OwnerType = "Investor"
	LocationType OwnerType = "Location"
	OfficeType   OwnerType = "Office"
)

type Address struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"userId"`
	OwnerID   uint   `json:"ownerId"`
	OwnerType string `json:"ownerType"`

	City    string `json:"city"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	ZipCode string `json:"zipCode"`
}
