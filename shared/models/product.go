package models

import (
	"fmt"
	"log"

	"github.com/lucsky/cuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID string `gorm:"primary_key"`
	Name string `json:"name"`
	CategoryID string `json:"category_id"`
	VendorID string `json:"vendor_id"`
	Hot bool `json:"hot"`
	Available int `json:"available"`
	BrandID string `json:"brand_id"`
}

type ProductDBService struct {
	DB *gorm.DB
}

type ProductDB interface {
	CreateProduct(userID string, vendorID string, name string, categoryID string, brandID string, available int) (*Product, error)
	GetProductsByVendor(vendorID string) ([]*Product, error)
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	if product.ID == "" {
		scope.SetColumn("ID", cuid.New())
	}

	if product.CategoryID == "" {
		scope.SetColumn("CategoryID", cuid.New())
	}

	if product.BrandID == "" {
		scope.SetColumn("BrandID", cuid.New())
	}

	return nil
}

func (service *ProductDBService) CreateProduct(userID string, vendorID string, name string, categoryID string, brandID string, available int) (*Product, error) {
	product := Product{
		VendorID:    vendorID,
		// CategoryID:  categoryID,
		// BrandID:     brandID,
		Name:        name,
		Available:   available,
	}

	result := service.DB.Create(&product)

	if result.Error != nil {
		log.Printf("An error occurred creating product %v", result.Error.Error())
		return &product, fmt.Errorf("an error occurred creating product %s", result.Error.Error())
	}

	return &product, nil
}

func (service *ProductDBService) GetProductsByVendor(vendorID string) ([]*Product, error) {
	var vendorProducts []*Product

	result := service.DB.Where("vendor_id = ?", vendorID).Find(&vendorProducts)

	if result.Error != nil {
		log.Printf("An error occurred getting all vendor products %v", result.Error.Error())
		return vendorProducts, fmt.Errorf("An error occurred getting all vendor products %s", result.Error.Error())
	}

	return vendorProducts, nil
}