package models

import "time"

type Note struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	UserID     uint `json:"userId"`
	LocationID uint `json:"locationId"`

	CreationDate *time.Time `json:"creationDate"`
	Content      string     `json:"content"`
}
