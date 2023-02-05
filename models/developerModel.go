package models

type Developer struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `json:"user_id"`

	Name string `gorm:"unique" json:"name"`
}
