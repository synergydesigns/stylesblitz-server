package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lucsky/cuid"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type ProductCategory struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
	ParentID    string
	AssetID     string
	VendorID    string
	ShopID      string
}

type ProductCategoryDBService struct {
	DB *gorm.DB
}

type ProductCategoryDB interface {
	GetAllCategories(vendorID, shopID string) ([]*ProductCategory, error)
	CreateCategory(vendorID, shopID, assetID, name, description, parentID string) (*ProductCategory, error)
	UpdateCategory(id, vendorID, shopID, assetID, name, description string) (*ProductCategory, error)
	DeleteCategory(id, vendorID string) (bool, error)
}

func (category *ProductCategory) BeforeCreate(scope *gorm.Scope) error {

	err := scope.SetColumn("ID", cuid.New())

	return err
}

func (service *ProductCategoryDBService) CreateCategory(vendorID, shopID, assetID, name, description, parentID string) (*ProductCategory, error) {
	category := ProductCategory{
		VendorID:    vendorID,
		Name:        name,
		Description: description,
		ParentID:    parentID,
		ShopID:      shopID,
	}

	if assetID != "" {
		category.AssetID = assetID
	}

	result := service.DB.Create(&category)

	if result.Error != nil {
		log.Printf("An error occurred creating category %v", result.Error.Error())

		if utils.HasRecord(result.Error) {
			return &category, fmt.Errorf("Category with name %s already exit", name)
		}

		if utils.ForeignKeyNotExist(result.Error) {
			return &category, fmt.Errorf("Vendor with id %s does not exit", vendorID)
		}

		return &category, fmt.Errorf("An error occurred creating category %v", result.Error)
	}

	return &category, nil
}

func (service *ProductCategoryDBService) GetAllCategories(vendorID, shopID string) ([]*ProductCategory, error) {
	var categories []*ProductCategory

	result := service.DB.Where("vendor_id = ? AND shop_id = ?", vendorID, shopID).Limit(20).Find(&categories)

	if result.Error != nil {
		log.Printf("An error occurred getting all categories %v", result.Error.Error())
		return categories, fmt.Errorf("An error occurred getting all categories %s", result.Error.Error())
	}

	return categories, nil
}

func (service *ProductCategoryDBService) UpdateCategory(id, vendorID, shopID, assetID, name, description string) (*ProductCategory, error) {
	category := ProductCategory{}
	value := make(map[string]interface{})

	if name != "" {
		value["name"] = name
	}

	if description != "" {
		value["description"] = description
	}

	result := service.DB.Model(&category).Where("id = ?", id).Where("vendor_id = ? AND shop_id = ?", vendorID, shopID).Updates(value)
	if result.Error != nil {
		log.Printf("An error occurred updating category %v", result.Error.Error())
		return &category, fmt.Errorf("An error occurred updating category %s", result.Error.Error())
	}

	result.First(&category, "id = ?", id)

	return &category, nil
}

func (service *ProductCategoryDBService) DeleteCategory(id, vendorID string) (bool, error) {
	category := ProductCategory{}

	result := service.DB.Delete(&category, "id = ? AND vendor_id = ?", id, vendorID)

	if result.Error != nil {
		log.Printf("An error occurred deleting category %v", result.Error.Error())
		return false, fmt.Errorf("An error occurred deleting category %s", result.Error.Error())
	}

	if (result.RowsAffected < 1) {
		return false, fmt.Errorf("An error occurred deleting category")
	}

	return true, nil
}
