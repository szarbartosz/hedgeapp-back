package initializers

import (
	"fmt"
	"main/models"
)

func SyncDB() {
	tables := []any{
		&models.User{},
		&models.Status{},
		&models.Address{},
		&models.Office{},
		&models.Investor{},
		&models.Location{},
		&models.Application{},
		&models.Note{},
	}

	for _, model := range tables {
		if err := DB.AutoMigrate(model); err != nil {
			panic(fmt.Sprintf("Failed to migrate %T: %s", model, err))
		}
	}
}
