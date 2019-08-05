package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service defines the service models for graphql
// for getting a single service
type Service struct {
	ID         uint64 `gorm:"primary_key"`
	Name       string
	Duration   int32
	Price      int32
	Status     bool
	Trend      int32
	VendorID   uint64 `json:"Vendor_id"`
	CategoryID uint64 `json:"category_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceDBService struct {
	DB *gorm.DB
}

type ServiceDB interface {
	GetServices(serviceName string, lat, long, radius float64) ([]*Service, error)
}

// ServiceQuery used for extracting user query
type ServiceQuery struct {
	Longitude bool
	Latitude  bool
}

// GetServices gets all services by query
func (service *ServiceDBService) GetServices(serviceName string, lat float64, long float64, radius float64) ([]*Service, error) {
	var services []*Service

	return services, nil
}
