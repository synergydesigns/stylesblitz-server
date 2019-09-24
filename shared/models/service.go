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
	GetServicesByVendor(vendorID string) ([]*Service, error)
	GetServicesByCategory(vendorID string, categoryID uint64) ([]*Service, error)
	CreateService(service ServiceInput) (*Service, error)
	UpdateService(id uint64, service ServiceInputUpdate) (*Service, error)
	DeleteService(id uint64) (bool, error)
	SearchService(lat *float64, lng *float64, name string, rating *SortRating, price *SortPrice) ([]*Service, error)
}

func (service *ServiceDBService) GetServices(serviceName string, lat float64, long float64, radius float64) ([]*Service, error) {
	var services []*Service

	return services, nil
}

func (service *ServiceDBService) GetServicesByVendor(vendorID string) ([]*Service, error) {
	var vendorServices []*Service

	result := service.DB.Where("vendor_id = ?", vendorID).Find(&vendorServices)

	if result.Error != nil {
		log.Printf("An error occurred getting all vendor services %v", result.Error.Error())
		return vendorServices, fmt.Errorf("An error occurred getting all vendor services %s", result.Error.Error())
	}

	return vendorServices, nil
}

func (service *ServiceDBService) GetServicesByCategory(vendorID string, categoryID uint64) ([]*Service, error) {
	var vendorServices []*Service

	result := service.DB.Where("vendor_id = ?", vendorID).Where("category_id = ?", categoryID).Find(&vendorServices)

	if result.Error != nil {
		log.Printf("An error occurred getting all category services %v", result.Error.Error())
		return vendorServices, fmt.Errorf("An error occurred getting all category services %s", result.Error.Error())
	}

	return vendorServices, nil
}

func (service *ServiceDBService) CreateService(serviceInput ServiceInput) (*Service, error) {
	newService := Service{
		Name:         serviceInput.Name,
		Duration:     uint(serviceInput.Duration),
		DurationType: serviceInput.DurationType.String(),
		CategoryID:   uint64(serviceInput.CategoryID),
		VendorID:     serviceInput.VendorID,
	}

	if serviceInput.Price != nil {
		newService.Price = *serviceInput.Price
	}

	if serviceInput.Trending != nil {
		newService.Trending = *serviceInput.Trending
	}

	result := service.DB.Create(&newService)

	if result.Error != nil {
		log.Printf("An error occurred creating service %v", result.Error.Error())
		if utils.HasRecord(result.Error) {
			return &newService, fmt.Errorf("service with name %s already exit", serviceInput.Name)
		}

		if utils.ForeignKeyNotExist(result.Error) {
			if utils.HasValue(result.Error.Error(), "vendor") {
				return &newService, fmt.Errorf("vendor with id %s does not exit", serviceInput.VendorID)
			}

			return &newService, fmt.Errorf("category with id %d does not exit", serviceInput.CategoryID)
		}

		return &newService, fmt.Errorf("an error occurred creating category %s", result.Error.Error())
	}

	return &newService, nil
}

// @TODO we need to decide how to handle updating categoryID for for vendor
// users should not be able to set service to a categoryID they did not create
// that should either be handle here or on the resolver level
func (service *ServiceDBService) UpdateService(id uint64, serviceInput ServiceInputUpdate) (*Service, error) {
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
		log.Printf("An error occurred updating service %v", result.Error.Error())
		return &vendorService, fmt.Errorf("An error occurred updating service %s", result.Error.Error())
	}

	result.First(&vendorService, "id = ?", id)

	return &vendorService, nil
}

func (service *ServiceDBService) DeleteService(id uint64) (bool, error) {
	result := service.DB.Delete(&Service{}, "id = ?", id)

	if result.Error != nil {
		log.Printf("An error occurred deleting service %v", result.Error.Error())
		return false, fmt.Errorf("An error occurred deleting service %s", result.Error.Error())
	}

	return true, nil
}

func (service *ServiceDBService) SearchService(
	lat *float64, lng *float64, name string, rating *SortRating, price *SortPrice,
) ([]*Service, error) {
	var allServices []*Service

	query := searchServiceQuery(lat, lng, name, rating, price)

	result := service.DB.Raw(query).Scan(&allServices)

	if result.Error != nil {
		log.Printf("An error getting all services %v", result.Error.Error())
		return allServices, fmt.Errorf("An error getting all services %s", result.Error.Error())
	}

	return allServices, nil
}
