package models

import (
	"fmt"

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

	sql := `SELECT DISTINCT
			p.id,
			p.name,
			p.description,
			p.about,
			p.phone,
			r.distance_unit * DEGREES(
				ACOS(
					COS(RADIANS(r.latpoint)) * COS(RADIANS(a.latitude)) * COS(RADIANS(r.longpoint) - RADIANS(a.longitude)) + SIN(RADIANS(r.latpoint)) * SIN(RADIANS(a.latitude))
				)
			) AS distance_in_km
		FROM
			provider as p
			LEFT JOIN service s on s.provider_id = p.id
			LEFT JOIN address a on a.provider_id = p.id
			LEFT JOIN (
				/* these are the query parameters */
				SELECT
					? AS latpoint,
					? AS longpoint,
					? AS radius,
					111.045 AS distance_unit
			) AS r ON 1 = 1
		WHERE
			match (s.name) against(? in natural language mode)
			AND a.latitude BETWEEN r.latpoint - (r.radius / r.distance_unit)
			AND r.latpoint + (r.radius / r.distance_unit)
			AND a.longitude BETWEEN r.longpoint - (
				r.radius / (r.distance_unit * COS(RADIANS(r.latpoint)))
			)
			AND r.longpoint + (
				r.radius / (r.distance_unit * COS(RADIANS(r.latpoint)))
			)
		GROUP BY
			p.name
		ORDER BY
			distance_in_km`

	rows, err := service.DB.Raw(sql, lat, long, radius, serviceName).Rows()

	if err != nil {
		return nil, fmt.Errorf("An error occurred getting services: %v", err.Error())
	}

	for rows.Next() {
		var provider Provider

		service.DB.ScanRows(rows, &provider)

		providers = append(providers, &provider)
	}

	return providers, nil
}
