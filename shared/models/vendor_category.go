package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type VendorCategory struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string
	Description string `json:"description"`
	VendorID    string `json:"vendor_id" gorm:"foreignkey:vendor_id"`
	Vendor      Vendor
}

type VendorCategoryDBService struct {
	DB *gorm.DB
}

type VendorCategoryDB interface {
	GetAllCategoriesByVendorID(vendorID string) ([]*VendorCategory, error)
	CreateCategory(vendorID, name, description string) (VendorCategory, error)
	UpdateCategory(id uint64, vendorID string, name, description *string) (VendorCategory, error)
	DeleteCategory(id uint64) (bool, error)
}

func (VendorCategory) TableName() string {
	return "categories"
}

func (category *VendorCategory) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("Name", strings.ToLower(category.Name))

	return err
}

func (service *VendorCategoryDBService) GetAllCategoriesByVendorID(vendorID string) ([]*VendorCategory, error) {
	var categories []*VendorCategory

	result := service.DB.Where("vendor_id = ?", vendorID).Limit(20).Find(&categories)

	if result.Error != nil {
		log.Printf("An error occurred getting all categories %v", result.Error.Error())
		return categories, fmt.Errorf("An error occurred getting all categories %s", result.Error.Error())
	}

	return categories, nil
}

func (service *VendorCategoryDBService) CreateCategory(vendorID, name, description string) (VendorCategory, error) {
	category := VendorCategory{
		VendorID:    vendorID,
		Name:        name,
		Description: description,
	}

	result := service.DB.Create(&category)

	if result.Error != nil {
		log.Printf("An error occurred creating category %v", result.Error.Error())

		if utils.HasRecord(result.Error) {
			return category, fmt.Errorf("Category with name %s already exit", name)
		}

		if utils.ForeignKeyNotExist(result.Error) {
			return category, fmt.Errorf("Vendor with id %s does not exit", vendorID)
		}

		return category, fmt.Errorf("An error occurred creating category %v", result.Error)
	}

	return category, nil
}

func (service *VendorCategoryDBService) UpdateCategory(id uint64, vendorID string, name, description *string) (VendorCategory, error) {
	category := VendorCategory{}
	value := make(map[string]interface{})

	if name != nil {
		value["name"] = *name
	}

	if description != nil {
		value["description"] = *description
	}

	result := service.DB.Model(&category).Where("id = ?", id).Where("vendor_id = ?", vendorID).Updates(value)
	if result.Error != nil {
		log.Printf("An error occurred updating category %v", result.Error.Error())
		return category, fmt.Errorf("An error occurred updating category %s", result.Error.Error())
	}

	result.First(&category, "id = ?", id)

	return category, nil
}

func (service *VendorCategoryDBService) DeleteCategory(id uint64) (bool, error) {
	category := VendorCategory{}

	result := service.DB.Delete(&category, "id = ?", id)

	if result.Error != nil {
		log.Printf("An error occurred deleting category %v", result.Error.Error())
		return false, fmt.Errorf("An error occurred deleting category %s", result.Error.Error())
	}

	return true, nil
}
