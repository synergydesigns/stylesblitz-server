package resolver

import (
	"context"
	"log"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type VendorCategoryResolver struct{}

func (query *queryResolver) GetAllCategories(ctx context.Context, vendorID *string) ([]*models.VendorCategory, error) {
	service := config.GetServices(ctx)

	categories, err := service.Datastore.VendorCategoryDB.GetAllCategoriesByVendorID(*vendorID)

	return categories, err
}

func (mutation mutationResolver) CreateVendorCategory(ctx context.Context, input models.VendorCategoryInput) (*models.VendorCategory, error) {
	service := config.GetServices(ctx)
	log.Println(input.Description)

	if input.Description == nil {
		input.Description = utils.StringToPointer("")
	}

	category, err := service.Datastore.VendorCategoryDB.CreateCategory(
		input.VendorID,
		input.Name,
		string(*input.Description),
	)

	return &category, err
}

func (mutation mutationResolver) UpdateVendorCategory(ctx context.Context, input models.VendorCategoryInputUpdate, categoryID int) (*models.VendorCategory, error) {
	service := config.GetServices(ctx)

	category, err := service.Datastore.VendorCategoryDB.UpdateCategory(
		uint64(categoryID),
		input.VendorID,
		input.Name,
		input.Description,
	)

	return &category, err
}

func (mutation mutationResolver) DeleteVendorCategory(ctx context.Context, categoryID int) (*bool, error) {
	service := config.GetServices(ctx)

	ok, err := service.Datastore.VendorCategoryDB.DeleteCategory(uint64(categoryID))

	return &ok, err
}
