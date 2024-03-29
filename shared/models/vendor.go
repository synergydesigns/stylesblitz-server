package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lucsky/cuid"
)

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

func (vendor *Vendor) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("Name", strings.ToLower(vendor.Name))

	if err != nil {
		return err
	}

	err = scope.SetColumn("Email", strings.ToLower(vendor.Email))

	if err != nil {
		return err
	}

	if vendor.ID == "" {
		err = scope.SetColumn("ID", cuid.New())
	}

	return err
}

func (service *VendorDbService) GetVendorsByServiceAndLocation(serviceName string, lat float64, long float64, radius float64) ([]*Vendor, error) {
	var Vendors []*Vendor

	return Vendors, nil
}
