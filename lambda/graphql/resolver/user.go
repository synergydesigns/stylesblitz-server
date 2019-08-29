package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	user, err := svc.Datastore.UserDB.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mutationResolver) CreateAccount(ctx context.Context, name *string) (*models.Asset, error) {
	panic(1)
}
