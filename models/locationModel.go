package models

import (
	"time"
)

type Location struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `json:"userId"`
	StatusID    uint        `json:"statusId"`
	InvestorID  uint        `json:"investorId"`
	OfficeID    uint        `json:"officeId"`
	Address     Address     `gorm:"polymorphic:Owner;" json:"address"`
	Application Application `json:"application"`
	Notes       []Note      `json:"notes"`

	Name              string     `gorm:"unique" json:"name"`
	IssueDate         *time.Time `json:"issueDate"`
	InspectionDate    *time.Time `json:"inspectionDate"`
	DecisionDate      *time.Time `json:"decisionDate"`
	DeforestationDate *time.Time `json:"deforestationDate"`
	PlantingDate      *time.Time `json:"plantingDate"`
	InspectionDone    bool       `json:"inspectionDone"`
	DeforestationDone bool       `json:"deforestationDone"`
	PlantingDone      bool       `json:"plantingDone"`
}
