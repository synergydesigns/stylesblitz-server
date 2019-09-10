package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
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

func (categoryService *queryResolver) GetAllVendorService(ctx context.Context, vendorID string) ([]*models.Service, error) {
	service := config.GetServices(ctx)

	createdServices, err := service.Datastore.ServiceDB.GetServicesByVendor(vendorID)

	return createdServices, err
}

func (categoryService *mutationResolver) CreateService(ctx context.Context, input models.ServiceInput) (*models.Service, error) {
	service := config.GetServices(ctx)

	createdService, err := service.Datastore.ServiceDB.CreateService(input)

	return createdService, err
}

func (categoryService *mutationResolver) UpdateService(ctx context.Context, input models.ServiceInputUpdate, serviceID int) (*models.Service, error) {
	service := config.GetServices(ctx)

	updatedService, err := service.Datastore.ServiceDB.UpdateService(uint64(serviceID), input)

	return updatedService, err
}

func (categoryService *mutationResolver) DeleteService(ctx context.Context, serviceID int) (*bool, error) {
	service := config.GetServices(ctx)

	ok, err := service.Datastore.ServiceDB.DeleteService(uint64(serviceID))

	return &ok, err
}
