package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

// ProviderResolver struct for resolving providers
type providerResolver struct{ *Resolver }

func (r *providerResolver) Addresses(ctx context.Context, obj *models.Provider) ([]*models.Address, error) {
	panic("not implemented")
}
func (r *providerResolver) Services(ctx context.Context, obj *models.Provider) ([]*models.Service, error) {
	panic("not implemented")
}

// GetProvidersByServiceAndLocation base on user query
func (r *queryResolver) GetProvidersByServiceAndLocation(ctx context.Context, name string, latitude *float64, longitude *float64, radius *float64, limit *int, page *int) ([]models.Provider, error) {
	var results []models.Provider

	return results, nil
}
