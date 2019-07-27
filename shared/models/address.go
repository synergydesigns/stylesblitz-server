package models

// Address defines the address models for graphql
// for getting a single address
type Address struct {
	ID        uint64 `gorm:"primary_key"`
	CountryID uint64
	StateID   uint64
	City      string
	Street    string
	Zipcode   int
	Longitude float64
	Latitude  float64
	Country   Country
	State     State
}
