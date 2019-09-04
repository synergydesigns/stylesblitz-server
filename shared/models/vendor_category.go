package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type VendorCategory struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string
	Description string
	VendorID    string `gorm:"foreignkey:vendor_id"`
	Vendor      Vendor
}

type VendorCategoryDBService struct {
	DB *gorm.DB
}

type VendorCategoryDB interface {
	GetAllCategoriesByVendorID(vendorID string) ([]VendorCategory, error)
}

func (VendorCategory) TableName() string {
	return "categories"
}

func (service *VendorCategoryDBService) GetAllCategoriesByVendorID(vendorID string) ([]VendorCategory, error) {
	var categories []VendorCategory

	result := service.DB.Where("vendor_id = ?", vendorID).Limit(20).Find(&categories)

	if result.Error != nil {
		log.Printf("An error occurred getting all categories %v", result.Error.Error())
		return categories, fmt.Errorf("An error occurred getting all categories %s", result.Error.Error())
	}

	return categories, nil
}
