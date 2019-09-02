package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type queryResolver struct{ *Resolver }

func (query *queryResolver) GetAllCategories(ctx context.Context, vendorID *string) ([]*models.VendorCategory, error) {
	panic(0)
}
