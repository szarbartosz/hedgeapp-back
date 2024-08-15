package models

import "time"

type CommonModel struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
