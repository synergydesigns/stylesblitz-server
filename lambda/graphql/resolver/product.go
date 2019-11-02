package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type productResolver struct{ *Resolver }

func (mutation mutationResolver) CreateProduct(ctx context.Context, input models.ProductInput) (*models.Product, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	product, err := service.Datastore.ProductDB.CreateProduct(
		user.ID,
		input.VendorID,
		input.Name,
		input.CategoryID,
		input.BrandID,
		input.Available,
	)

	return product, err
}

func (query *queryResolver) GetProductsByVendor(ctx context.Context, vendorID string) ([]*models.Product, error) {
	service := config.GetServices(ctx)

	products, err := service.Datastore.ProductDB.GetProductsByVendor(vendorID)

	return products, err
}
