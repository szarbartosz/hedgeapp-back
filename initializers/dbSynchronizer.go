package initializers

import "main/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
