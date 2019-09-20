package model

import (
	"github.com/jinzhu/gorm"
)

type ServiceCart struct {
	ID int `gorm:"primary_key"`
	VendorID string `json:"vendor_id"`
	ServiceID string `json:"service_id"`
	CartID string `json:"cart_id"`
	Quantity int `gorm:"default:1"`
}

type ServiceCartDBService struct {
	DB *gorm.DB
}
