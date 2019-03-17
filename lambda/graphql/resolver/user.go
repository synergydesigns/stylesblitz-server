package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

// userResolver struct for resolving use
type userResolver struct{ *Resolver }

// User user
func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	user, err := svc.Datastore.UserDB.GetUserByID(uint64(id))
	if err != nil {
		return nil, err
	}

	return user, nil
}
