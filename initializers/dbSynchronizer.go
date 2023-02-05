package initializers

import "main/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.Developer{})
	DB.AutoMigrate(&models.Location{})
}
