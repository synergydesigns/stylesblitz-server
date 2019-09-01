package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Vendor defines the Vendor models for graphql
// for getting a single Vendor
type Vendor struct {
	ID           string `gorm:"primary_key"`
	Name         string
	Description  string
	Phone        JSON `sql:"type:json"`
	User         *User
	ProfileImage string
	UserID       string
	Email        string
	Verified     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Phone struct {
	Code  uint32
	Phone uint64
}

type VendorDbService struct {
	DB *gorm.DB
}

type VendorDB interface {
	GetVendorsByServiceAndLocation(serviceName string, lat, long, radius float64) ([]*Vendor, error)
}

// GetVendorsByServiceAndLocation gets all services by query
func (service *VendorDbService) GetVendorsByServiceAndLocation(serviceName string, lat float64, long float64, radius float64) ([]*Vendor, error) {
	var Vendors []*Vendor

	return Vendors, nil
}
