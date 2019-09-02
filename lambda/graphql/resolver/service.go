package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type serviceResolver struct{}

func (service *serviceResolver) Duration(ctx context.Context, obj *models.Service) (*int, error) {
	duration := int(obj.Duration)
	return &duration, nil
}

func (service *serviceResolver) Price(ctx context.Context, obj *models.Service) (*int, error) {
	price := int(obj.Price)
	return &price, nil
}
