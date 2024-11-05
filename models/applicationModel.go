package models

import "time"

type Application struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	UserID     uint `json:"userId"`
	LocationID uint `json:"locationId"`

	Signature          string     `json:"signature"`
	IsCommercial       bool       `json:"isCommercial"`
	DeforestationCause string     `json:"deforestationCause"`
	DeforestationDate  *time.Time `json:"deforestationDate"`
	PlantingDate       *time.Time `json:"plantingDate"`
	PlantingPlace      string     `json:"plantingPlace"`
	Species            string     `json:"species"`
}
