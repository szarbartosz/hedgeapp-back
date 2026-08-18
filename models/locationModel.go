package models

import (
	"time"
)

type Location struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `json:"userId"`
	StatusID    uint        `json:"statusId" binding:"required"`
	Status      Status      `gorm:"foreignKey:StatusID" json:"status"`
	OfficeID    uint        `json:"officeId" binding:"required"`
	Office      Office      `gorm:"foreignKey:OfficeID" json:"office"`
	InvestorID  uint        `json:"investorId" binding:"required"`
	Investor    Investor    `gorm:"foreignKey:InvestorID" json:"investor"`
	AddressID   uint        `json:"addressId"`
	Address     Address     `gorm:"foreignKey:AddressID" json:"address"`
	Application Application `json:"application"`
	Notes       []Note      `json:"notes"`

	Name              string     `json:"name"`
	IssueDate         *time.Time `json:"issueDate"`
	InspectionDate    *time.Time `json:"inspectionDate"`
	DecisionDate      *time.Time `json:"decisionDate"`
	DeforestationDate *time.Time `json:"deforestationDate"`
	PlantingDate      *time.Time `json:"plantingDate"`
	InspectionDone    bool       `json:"inspectionDone"`
	DeforestationDone bool       `json:"deforestationDone"`
	PlantingDone      bool       `json:"plantingDone"`
}
