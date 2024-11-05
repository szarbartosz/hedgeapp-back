package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("No .env file found - that's okay if running on docker.")
		} else {
			log.Panic("Error loading .env file!")
		}
	}
}
