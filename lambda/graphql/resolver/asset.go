package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

func (r *queryResolver) GetAsset(ctx context.Context, id string) (*models.Asset, error) {
	var asset models.Asset
	return &asset, nil
}
