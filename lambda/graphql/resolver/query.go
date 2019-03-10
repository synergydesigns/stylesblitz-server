package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAddress(ctx context.Context, id int) (*models.Address, error) {
	panic("not implemented")
}
