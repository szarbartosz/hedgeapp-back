package models

type Investor struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"userId"`
	Locations []Location `json:"locations"`
	Address   Address    `gorm:"polymorphic:Owner;" json:"address"`

	Name          string `gorm:"unique" json:"name"`
	ContactPerson string `json:"contactPerson"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Nip           string `json:"nip"`
	Regon         string `json:"regon"`
}
