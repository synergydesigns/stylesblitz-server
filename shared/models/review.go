package models

import (
	"github.com/jinzhu/gorm"
)

// Review defines the review models for graphql
// for getting a single review
type Review struct {
	gorm.Model
	UserID uint
	ShopID uint
	Rating string
	Review string
}
