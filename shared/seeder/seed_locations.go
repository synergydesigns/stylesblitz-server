package seeder

import (
	"encoding/json"
	"fmt"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

var seeder = New()

func seedCountries() {
	var countries []models.Country
	// var
	json.Unmarshal(
		getData("countries"),
		&countries,
	)

	for _, country := range countries {
		func(country models.Country) {
			fmt.Println(country)
			seeder.DB.Table("countries").Create(&country)
		}(country)
	}
}

func SeedLocations() {
	seedCountries()
}
