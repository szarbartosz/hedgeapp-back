package models

type Office struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	AddressID uint    `json:"addressId"`
	Address   Address `gorm:"foreignKey:AddressID" json:"address"`

	Name string `json:"name"`
}
