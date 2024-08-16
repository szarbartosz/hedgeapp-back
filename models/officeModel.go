package models

type Office struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Address Address `gorm:"polymorphic:Owner;" json:"address"`

	Name string `json:"name"`
}
