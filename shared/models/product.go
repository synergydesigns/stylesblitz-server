package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID string `gorm:"primary_key"`
	CategoryID string `json:"category_id"`
	Hot bool `json:"hot"`
	Available int `json:"available"`
	BrandID string `json:"brand_id"`
}

type ProductDBService struct {
	DB *gorm.DB
}
