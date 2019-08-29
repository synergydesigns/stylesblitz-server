package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

func (r *queryResolver) GetAsset(ctx context.Context, id string) (*models.Asset, error) {
	var asset models.Asset

	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	asset, err := svc.Datastore.AssetDB.GetAsset(id)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (m *mutationResolver) CreatePresignedURL(ctx context.Context, input []*models.AssetInput) (*string, error) {
	panic(1)
}
