package db

import (
	"github.com/jinzhu/gorm"
)

// Service defines the service models for graphql
// for getting a single service
type Service struct {
	gorm.Model
	Name     string
	Duration int
	Price    int
	Status   string
	Trend    string
	ShopID   uint
}
