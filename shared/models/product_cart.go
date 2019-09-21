package models

import (
	"github.com/jinzhu/gorm"
)

type ProductCart struct {
	ID int `gorm:"primary_key"`
	VendorID string `json:"vendor_id"`
	ProductID string `json:"product_id"`
	ShopID string `json:"shop_id"`
	CartID string `json:"cart_id"`
	Quantity int `gorm:"default:1"`
}

type ProductCartDBService struct {
	DB *gorm.DB
}
