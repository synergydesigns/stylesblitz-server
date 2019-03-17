package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type serviceResolver struct{ *Resolver }

// Services base on user query
func (r *queryResolver) Services(ctx context.Context, name string, latitude *float64, longitude *float64, radius *float64, limit *int, page *int) ([]models.Service, error) {
	var s []models.Service
	return s, nil
}
