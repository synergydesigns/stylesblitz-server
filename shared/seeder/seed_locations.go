package seeder

import (
	"encoding/json"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

var seeder = New()

type stateData struct {
	Name      string
	StateCode string
	Country   string
	CountryID uint64
	Longitude float32
	Latitude  float32
}

type cityData struct {
	Name      string
	Longitude float32
	Latitude  float32
	State     string
}

// SeedCountries inserts all countries fields in the countries.json file
func SeedCountries() {
	var countries []models.Country

	json.Unmarshal(
		getData("countries"),
		&countries,
	)

	for _, country := range countries {
		func(country models.Country) {
			seeder.DB.Table("countries").Create(&country)
		}(country)
	}
}

// SeedStates inserts all states fields in the states.json file
func SeedStates() {
	var states []stateData
	var countries []models.Country

	seeder.DB.Table("countries").Select("*").Scan(&countries)

	json.Unmarshal(
		getData("states"),
		&states,
	)

	for _, country := range countries {
		for _, state := range states {
			if country.Name == state.Country {
				state.CountryID = country.ID
				newState := models.State{
					Name:      state.Name,
					StateCode: state.StateCode,
					CountryID: country.ID,
					Longitude: state.Longitude,
					Latitude:  state.Latitude,
				}

				seeder.DB.Table("states").Create(&newState)
			}
		}
	}
}

// SeedCities inserts all cities fields in the citites.json file
func SeedCities() {
	var cities []cityData
	var states []models.State

	seeder.DB.Table("states").Select("*").Scan(&states)

	json.Unmarshal(
		getData("cities"),
		&cities,
	)

	for _, state := range states {
		for _, city := range cities {
			stateName := state.Name + state.StateCode
			if stateName == city.State {
				func(city cityData, state models.State) {
					newCity := models.City{
						Name:      city.Name,
						CountryID: state.CountryID,
						StateID:   state.ID,
						Longitude: city.Longitude,
						Latitude:  city.Latitude,
					}
					seeder.DB.Create(&newCity)
				}(city, state)
			}
		}

	}
}

// SeedLocations seeds all default locations countries -> states -> cities
func SeedLocations() {
	SeedCountries()
	SeedStates()
	SeedCities()
}
