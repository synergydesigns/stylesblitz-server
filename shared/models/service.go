package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type Service struct {
	ID           uint64 `gorm:"primary_key"`
	Name         string
	Duration     uint
	DurationType string
	Price        float64
	Trending     bool
	VendorID     string `json:"Vendor_id"`
	CategoryID   uint64 `json:"category_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ServiceDBService struct {
	DB *gorm.DB
}

type ServiceDB interface {
	GetServices(serviceName string, lat, long, radius float64) ([]*Service, error)
	GetServicesByVendor(vendorID string) ([]Service, error)
	GetServicesByCategory(vendorID, categoryID string) ([]Service, error)
	CreateService(service ServiceInput) (Service, error)
	UpdateService(id uint64, service ServiceInputUpdate) (bool, error)
}

func (service *ServiceDBService) GetServices(serviceName string, lat float64, long float64, radius float64) ([]*Service, error) {
	var services []*Service

	return services, nil
}

func (service *ServiceDBService) GetServicesByVendor(vendorID string) ([]Service, error) {
	var vendorServices []Service

	result := service.DB.Where("vendor_id = ?", vendorID).Find(&vendorServices)

	if result.Error != nil {
		log.Printf("An error occurred getting all vendor services %v", result.Error.Error())
		return vendorServices, fmt.Errorf("An error occurred getting all vendor services %s", result.Error.Error())
	}

	return vendorServices, nil
}

func (service *ServiceDBService) GetServicesByCategory(vendorID, categoryID string) ([]Service, error) {
	var vendorServices []Service

	result := service.DB.Where("vendor_id = ?", categoryID).Where("category_id = ?", categoryID).Find(&vendorServices)

	if result.Error != nil {
		log.Printf("An error occurred getting all vendor services %v", result.Error.Error())
		return vendorServices, fmt.Errorf("An error occurred getting all vendor services %s", result.Error.Error())
	}

	return vendorServices, nil
}

func (service *ServiceDBService) CreateService(serviceInput ServiceInput) (Service, error) {
	newService := Service{
		Name:         serviceInput.Name,
		Price:        *serviceInput.Price,
		Duration:     uint(serviceInput.Duration),
		DurationType: serviceInput.DurationType.String(),
		Trending:     *serviceInput.Trending,
		CategoryID:   uint64(serviceInput.CategoryID),
		VendorID:     serviceInput.VendorID,
	}

	result := service.DB.Create(&newService)

	if result.Error != nil {
		log.Printf("An error occurred creating category %v", result.Error.Error())

		if utils.HasRecord(result.Error) {
			return newService, fmt.Errorf("Service with name %d already exit", serviceInput.CategoryID)
		}

		if utils.ForeignKeyNotExist(result.Error) {
			return newService, fmt.Errorf("Vendor with id %s does not exit", serviceInput.VendorID)
		}

		return newService, fmt.Errorf("An error occurred creating category %v", result.Error)
	}

	return newService, nil
}

func (service *ServiceDBService) UpdateService(id uint64, serviceInput ServiceInputUpdate) (bool, error) {
	vendorService := Service{}
	fields := make(map[string]interface{})

	updateData := utils.StructToInterface(serviceInput)

	for key, value := range updateData {
		if value != nil {
			fields[key] = value
		}
	}

	result := service.DB.Model(&vendorService).Where("id = ?", id).Updates(fields)

	if result.Error != nil {
		log.Printf("An error occurred updating category %v", result.Error.Error())
		return false, fmt.Errorf("An error occurred updating category %s", result.Error.Error())
	}

	return true, nil
}
