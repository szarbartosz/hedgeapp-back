package models

type Investor struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"userId"`
	Address   Address    `json:"address"`
	Locations []Location `json:"locations"`

	Name string `gorm:"unique" json:"name"`
}
