package models

type Office struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `json:"userId"`

	Name string `json:"name"`
}
