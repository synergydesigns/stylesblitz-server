package models

import (
	"github.com/jinzhu/gorm"
)

// Address defines the address models for graphql
// for getting a single address
type Address struct {
	gorm.Model
	ProviderID uint `json:"provider_id"`
	Country    string
	State      string
	City       string
	Zipcode    string
	Longitude  float64
	Latitude   float64
}
