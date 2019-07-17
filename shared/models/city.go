package models

type City struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	StateID   uint64
	CountryID uint64
	Longitude float32
	Latitude  float32
}
