package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type productCategoryResolver struct {*Resolver}

func (mutation mutationResolver) CreateProductCategory(ctx context.Context, input models.ProductCategoryInput) (*models.ProductCategory, error) {
	service := config.GetServices(ctx)
	category, err := service.Datastore.ProductCategoryDB.CreateCategory(
		input.VendorID,
		input.ShopID,
		input.AssetID,
		input.Name,
		input.Description,
		input.ParentID,
	)

	return category, err
}

func (query *queryResolver) GetAllProductCategories(ctx context.Context, vendorID, shopID string) ([]*models.ProductCategory, error) {
	service := config.GetServices(ctx)
	
	categories, err := service.Datastore.ProductCategoryDB.GetAllCategories(vendorID, shopID)

	return categories, err
}

func (mutation mutationResolver) UpdateProductCategory(ctx context.Context, input models.ProductCategoryUpdateInput) (*models.ProductCategory, error) {
  service := config.GetServices(ctx)
	
	category, err := service.Datastore.ProductCategoryDB.UpdateCategory(
		input.ID,
		input.VendorID,
		input.ShopID,
		input.AssetID,
		input.Name,
		input.Description,
	)

	return category, err
}

func (mutation mutationResolver) DeleteProductCategory(ctx context.Context, id, vendorID string) (*bool, error) {
	service := config.GetServices(ctx)
	
	ok, err := service.Datastore.ProductCategoryDB.DeleteCategory(id, vendorID)

	return &ok, err
}
