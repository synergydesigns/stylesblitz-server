package models

import (
	"github.com/jinzhu/gorm"
)

// Category defines the category models for graphql
// for getting a single category
type Category struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	ProviderID  uint `json:"provider_id"`
}
