package initializers

import "main/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Note{})
	DB.AutoMigrate(&models.Office{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.Application{})
	DB.AutoMigrate(&models.Investor{})
	DB.AutoMigrate(&models.Location{})
}
