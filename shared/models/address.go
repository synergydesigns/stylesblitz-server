package models

// Address defines the address models for graphql
// for getting a single address
type Address struct {
	ID         uint64 `gorm:"primary_key"`
	ProviderID uint64 `json:"provider_id"`
	Country    string
	State      string
	City       string
	Zipcode    string
	Longitude  float64
	Latitude   float64
}
