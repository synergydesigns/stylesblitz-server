package models

import (
	"github.com/jinzhu/gorm"
)

// Provider defines the provider models for graphql
// for getting a single provider
type Provider struct {
	gorm.Model
	Name        string
	Description string
	About       string
	Phone       uint
	User        User
	Address     Address
	Opening     Opening
}
