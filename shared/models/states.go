package models

type State struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	StateCode string
	CountryID uint64 `json:"countryId"`
	Longitude float32
	Latitude  float32
}
