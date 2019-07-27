package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAddress(ctx context.Context, id int) (*models.Address, error) {
	panic("not implemented")
}

func (r *queryResolver) GetVendorsByServiceAndLocation(ctx context.Context, name string, latitude *float64, longitude *float64, radius *float64, limit *int, page *int) ([]models.Vendor, error) {
	var results []models.Vendor

	return results, nil
}
