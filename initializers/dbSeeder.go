package initializers

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"main/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:embed seeds/*.json
var seedFS embed.FS

var seededTables = []string{"statuses", "offices", "addresses"}

func SeedDB() {
	LoadStatusesFromJSON("seeds/statuses.json")
	LoadOfficesFromJSON("seeds/offices.json")
	syncSequences()
}

func upsertByID(value any, columns []string) *gorm.DB {
	return DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Omit(clause.Associations).Create(value)
}

func readSeed(filename string, dest any) error {
	data, err := seedFS.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("reading seed file: %w", err)
	}

	if err := json.Unmarshal(data, dest); err != nil {
		return fmt.Errorf("unmarshalling seed file: %w", err)
	}

	return nil
}

func LoadStatusesFromJSON(filename string) {
	var statuses []models.Status

	if err := readSeed(filename, &statuses); err != nil {
		log.Panicf("%s: %s", filename, err)
	}

	for _, status := range statuses {
		if result := upsertByID(&status, []string{"name"}); result.Error != nil {
			log.Printf("Error seeding status [%d] %s: %s", status.ID, status.Name, result.Error)
		} else {
			log.Printf("Seeded status [%d] %s", status.ID, status.Name)
		}
	}
}

func LoadOfficesFromJSON(filename string) {
	var offices []models.Office

	if err := readSeed(filename, &offices); err != nil {
		log.Panicf("%s: %s", filename, err)
	}

	for _, office := range offices {
		if result := upsertByID(&office.Address, []string{"city", "street", "number", "zip_code"}); result.Error != nil {
			log.Printf("Error seeding address for office [%d] %s: %s", office.ID, office.Name, result.Error)
			continue
		}

		office.AddressID = office.Address.ID

		if result := upsertByID(&office, []string{"name", "address_id"}); result.Error != nil {
			log.Printf("Error seeding office [%d] %s: %s", office.ID, office.Name, result.Error)
		} else {
			log.Printf("Seeded office [%d] %s", office.ID, office.Name)
		}
	}
}

func syncSequences() {
	for _, table := range seededTables {
		query := fmt.Sprintf(
			"SELECT setval(pg_get_serial_sequence('%s', 'id'), COALESCE((SELECT MAX(id) FROM %s), 0) + 1, false)",
			table, table,
		)

		if err := DB.Exec(query).Error; err != nil {
			log.Printf("Error resetting id sequence for %s: %s", table, err)
		}
	}
}
