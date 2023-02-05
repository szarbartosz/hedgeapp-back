package models

import (
	"time"
)

type Location struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	UserID      uint `json:"user_id"`
	StatusID    uint `json:"status_id"`
	DeveloperID uint `json:"developer_id"`

	Name string `gorm:"unique" json:"name"`
	// ISO 8601 format required for dates
	IssueDate         *time.Time `json:"issue_date"`
	InspectionDate    *time.Time `json:"inspection_date"`
	DeforestationDate *time.Time `json:"deforestation_date"`
	PlantingDate      *time.Time `json:"planting_date"`
}
