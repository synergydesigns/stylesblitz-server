package models

import (
	"fmt"
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
	VendorID uint64 `json:"Vendor_id"`
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

	sql := `SELECT
		*,
		p.distance_unit * DEGREES(
			ACOS(
				COS(RADIANS(p.latpoint)) * COS(RADIANS(a.latitude)) * COS(RADIANS(p.longpoint) - RADIANS(a.longitude)) + SIN(RADIANS(p.latpoint)) * SIN(RADIANS(a.latitude))
			)
		) AS distance_in_km
	FROM
		service AS s
		JOIN address a on s.Vendor_id = a.Vendor_id
		JOIN (
			SELECT
				? AS latpoint,
				? AS longpoint,
				? AS radius,
				111.045 AS distance_unit
		) AS p ON 1 = 1
	WHERE
		a.latitude BETWEEN p.latpoint - (p.radius / p.distance_unit)
		AND p.latpoint + (p.radius / p.distance_unit)
		AND a.longitude BETWEEN p.longpoint - (
			p.radius / (p.distance_unit * COS(RADIANS(p.latpoint)))
		)
		AND p.longpoint + (
			p.radius / (p.distance_unit * COS(RADIANS(p.latpoint)))
		)
	ORDER BY
		distance_in_km
	LIMIT
		15`

	rows, err := service.DB.Raw(sql, lat, long, radius).Rows()

	if err != nil {
		return nil, fmt.Errorf("An error occurred getting services: %v", err.Error())
	}

	for rows.Next() {
		var svc Service

		service.DB.ScanRows(rows, &svc)

		services = append(services, &svc)
	}

	return services, nil
}
