package models

import (
	"github.com/jinzhu/gorm"
)

// Tag defines the tag models for graphql
// for getting a single tag
type Tag struct {
	gorm.Model
	Name   string
	ShopID uint
}
