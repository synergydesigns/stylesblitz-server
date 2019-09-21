package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Shop struct {
	ID string `gorm:"primary_key"`
	Name string `json:"name"`
	VendorID int `json:"vendor_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShopDBService struct {
	DB *gorm.DB
}
