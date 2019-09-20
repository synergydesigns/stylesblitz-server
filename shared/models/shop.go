package model

import (
	"github.com/jinzhu/gorm"
)

type Shop struct {
	ID string `gorm:"primary_key"`
	Name string `json:"name"`
	VendorID int `json:"vendor_id"`
	CreatedAt bool `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type ShopDBService struct {
	DB *gorm.DB
}
