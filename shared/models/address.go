package models

import (
	"github.com/jinzhu/gorm"
)

// Address defines the address models for graphql
// for getting a single address
type Address struct {
	gorm.Model
	ShopID  uint
	Country string
	State   string
	City    string
	ZipCode string
}
