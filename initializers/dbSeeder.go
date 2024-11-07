package initializers

import (
	"encoding/json"
	"log"
	"main/models"
	"os"
)

func SeedDB() {
	LoadStatusesFromJSON("./seeds/statuses.json")
	LoadOfficesFromJSON("./seeds/offices.json")
}

func LoadStatusesFromJSON(filename string) {
	var statuses []models.Status

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("Error reading JSON file: %s", err)
	}

	err = json.Unmarshal(data, &statuses)
	if err != nil {
		log.Panicf("Error unmarshalling JSON: %s", err)
	}

	for _, status := range statuses {
		if result := DB.FirstOrCreate(&status, models.Status{ID: status.ID}); result.Error != nil {
			log.Printf("Error inserting/updating status [%d] %s: %s", status.ID, status.Name, result.Error)
		} else {
			log.Printf("Successfully inserted/updated status [%d] %s", status.ID, status.Name)
		}
	}
}

func LoadAddressesFromJSON(filename string) {
	var addresses []models.Address

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("Error reading JSON file: %s", err)
	}

	err = json.Unmarshal(data, &addresses)
	if err != nil {
		log.Panicf("Error unmarshalling JSON: %s", err)
	}

	for _, address := range addresses {
		if result := DB.FirstOrCreate(&address, models.Address{ID: address.ID}); result.Error != nil {
			log.Printf("Error inserting/updating address [%d] %s: %s", address.ID, address.City, result.Error)
		} else {
			log.Printf("Successfully inserted/updated address [%d] %s", address.ID, address.City)
		}
	}
}

func LoadOfficesFromJSON(filename string) {
	var offices []models.Office

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("Error reading JSON file: %s", err)
	}

	err = json.Unmarshal(data, &offices)
	if err != nil {
		log.Panicf("Error unmarshalling JSON: %s", err)
	}

	for _, office := range offices {
		if result := DB.FirstOrCreate(&office, models.Office{ID: office.ID}); result.Error != nil {
			log.Printf("Error inserting/updating office [%d] %s: %s", office.ID, office.Name, result.Error)
		} else {
			log.Printf("Successfully inserted/updated office [%d] %s", office.ID, office.Name)
		}
	}
}
