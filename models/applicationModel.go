package models

type Application struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	UserID     uint `json:"userId"`
	LocationID uint `json:"locationId"`

	Signature          string `json:"signature"`
	IsCommercial       string `json:"isCommercial"`
	DeforestationCause string `json:"deforestationCause"`
	DeforestationDate  string `json:"deforestationDate"`
	PlantingDate       string `json:"plantingDate"`
	PlantingSite       string `json:"plantingSite"`
	Species            string `json:"species"`
}
