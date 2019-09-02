package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Service struct {
	ID           uint64 `gorm:"primary_key"`
	Name         string
	Duration     uint
	DurationType string
	Price        uint
	Status       bool
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
}

func (service *ServiceDBService) GetServices(serviceName string, lat float64, long float64, radius float64) ([]*Service, error) {
	var services []*Service

	return services, nil
}
