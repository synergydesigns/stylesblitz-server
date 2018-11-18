package resolver

import (
	"context"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

// ProviderResolver struct for resolving providers
type ProviderResolver struct {
	provider *models.Provider
}

// ID provider id
func (r *ProviderResolver) ID() graphql.ID {
	ID := strconv.Itoa(int(r.provider.ID))

	return graphql.ID(ID)
}

// Name of service provider
func (r *ProviderResolver) Name() *string {
	return &r.provider.Name
}

// Description about service provider
func (r *ProviderResolver) Description() *string {
	return &r.provider.Description
}

// About provider name
func (r *ProviderResolver) About() *string {
	return &r.provider.About
}

// Phone provider trend
func (r *ProviderResolver) Phone() *string {
	return &r.provider.Phone
}

// User that provides the service
func (r *ProviderResolver) User() *userResolver {
	return &userResolver{&models.User{}}
}

// Addresses that provides the service
func (r *ProviderResolver) Addresses() *[]*addressResolver {
	var addresses []*addressResolver

	addresses = append(addresses, &addressResolver{&models.Address{}})
	return &addresses
}

// Services that provides the service
func (r *ProviderResolver) Services() *[]*serviceResolver {
	var services []*serviceResolver

	services = append(services, &serviceResolver{&models.Service{}})
	return &services
}

type getServiceQuery struct {
	Longitude, Latitude, Radius *float64
	Name                        string
	Limit, Page                 *int32
}

// GetProvidersByServiceAndLocation base on user query
func (r *Resolver) GetProvidersByServiceAndLocation(ctx context.Context, args getServiceQuery) (*[]*ProviderResolver, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)

	// set default limit and offset
	var limit int32
	var offset int32 = 1
	var radius float64 = 1

	if args.Limit == nil || (args.Limit != nil && *args.Limit == 0) {
		limit = config.ServicesLimit
	} else {
		limit = *args.Limit
	}

	if args.Page != nil && *args.Page > 0 {
		offset = *args.Page
	}

	if args.Radius != nil {
		radius = *args.Radius
	}

	offset = limit * (offset - 1)

	providers, err := svc.Datastore.ProviderDB.GetProvidersByServiceAndLocation(
		args.Name,
		*args.Latitude,
		*args.Longitude,
		radius,
	)
	if err != nil {
		return nil, err
	}

	var results []*ProviderResolver

	for _, provider := range providers {
		results = append(results, &ProviderResolver{provider})
	}
	return &results, nil
}
