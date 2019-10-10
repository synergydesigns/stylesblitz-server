package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type CartResolver struct{*Resolver}
func (mutation mutationResolver) CreateProductCart(ctx context.Context, input models.ProductCartInput) (*models.ProductCart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	productCart, err := service.Datastore.ProductCartDB.CreateProductCart(
		user.ID,
		input.VendorID,
		input.ProductID,
		input.Quantity,
	)

	return productCart, err
}

func (mutation mutationResolver) UpdateProductCart(ctx context.Context, input models.CartUpdateInput) (*models.ProductCart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	productCart, err := service.Datastore.ProductCartDB.UpdateProductCart(
		user.ID,
		input.CartID,
		input.Quantity,
	)

	return productCart, err
}

func (mutation mutationResolver) DeleteProductCart(ctx context.Context, cartID string) (*bool, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	ok, err := service.Datastore.ProductCartDB.DeleteProductCart(
		user.ID,
		cartID,
	)

	return &ok, err
}

func (query *queryResolver) GetProductsCart(ctx context.Context, userID *string) ([]*models.ProductCart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	productsCart, err := service.Datastore.ProductCartDB.GetProductsCart(
		user.ID,
	)

	return productsCart, err
}


func (mutation mutationResolver) CreateServiceCart(ctx context.Context, input models.ServiceCartInput) (*models.ServiceCart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	serviceCart, err := service.Datastore.ServiceCartDB.CreateServiceCart(
		user.ID,
		input.VendorID,
		input.ServiceID,
		input.Quantity,
	)

	return serviceCart, err
}

func (mutation mutationResolver) UpdateServiceCart(ctx context.Context, input models.CartUpdateInput) (*models.ServiceCart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	serviceCart, err := service.Datastore.ServiceCartDB.UpdateServiceCart(
		user.ID,
		input.CartID,
		input.Quantity,
	)

	return serviceCart, err
}

func (mutation mutationResolver) DeleteServiceCart(ctx context.Context, cartID string) (*bool, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	ok, err := service.Datastore.ServiceCartDB.DeleteServiceCart(
		user.ID,
		cartID,
	)

	return &ok, err
}

func (query *queryResolver) GetServicesCart(ctx context.Context, userID *string) ([]*models.ServiceCart, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	servicesCart, err := service.Datastore.ServiceCartDB.GetServicesCart(
		user.ID,
	)

	return servicesCart, err
}

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