package models

import (
	"github.com/jinzhu/gorm"
)

// Opening defines the opening models for graphql
// for getting a single opening
type Opening struct {
	gorm.Model
	VendorID uint
	From       string
	To         string
	Day        string
}
