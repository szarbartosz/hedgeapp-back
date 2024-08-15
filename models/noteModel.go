package models

type Note struct {
	CommonModel
	ID         uint `gorm:"primaryKey" json:"id"`
	UserID     uint `json:"userId"`
	LocationID uint `json:"locationId"`

	Content string `json:"content"`
}
