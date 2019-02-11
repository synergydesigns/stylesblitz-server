package models

import (
	"github.com/jinzhu/gorm"
)

// Provider defines the provider models for graphql
// for getting a single provider
type Provider struct {
	gorm.Model
	Name        string
	Description string
	About       string
	Phone       string
	User        User
	Addresses   []Address
	Opening     Opening
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
}

type ProviderDbService struct {
	DB *gorm.DB
}

type ProviderDB interface {
	GetProvidersByServiceAndLocation(serviceName string, lat, long, radius float64) ([]*Provider, error)
}

// GetProvidersByServiceAndLocation gets all services by query
func (service *ProviderDbService) GetProvidersByServiceAndLocation(serviceName string, lat float64, long float64, radius float64) ([]*Provider, error) {
	var providers []*Provider

	return providers, nil
}
