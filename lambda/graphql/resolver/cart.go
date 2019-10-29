package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type CartResolver struct{ *Resolver }

func (mutation mutationResolver) CreateCart(ctx context.Context, input models.CartInput) (*models.Cart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	cart, err := service.Datastore.CartDB.CreateCart(
		user.ID,
		input.VendorID,
		input.Type,
		input.TypeID,
		input.Quantity,
	)

	return cart, err
}

func (query *queryResolver) GetAllCarts(ctx context.Context, userID *string) ([]*models.Cart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	cart, err := service.Datastore.CartDB.GetAllCarts(
		user.ID,
	)

	return cart, err
}

func (mutation mutationResolver) UpdateCart(ctx context.Context, input models.CartUpdateInput) (*models.Cart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	cart, err := service.Datastore.CartDB.UpdateCart(
		user.ID,
		input.CartID,
		input.Quantity,
		input.Type,
		input.TypeID,
	)

	return cart, err
}

func (mutation mutationResolver) DeleteCart(ctx context.Context, cartID string) (*bool, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	ok, err := service.Datastore.CartDB.DeleteCart(
		user.ID,
		cartID,
	)

	return &ok, err
}
