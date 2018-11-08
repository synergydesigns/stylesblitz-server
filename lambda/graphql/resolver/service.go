package resolver

import (
	"context"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/config"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
	service "gitlab.com/synergy-designs/style-blitz/shared/services"
)

// ServiceResolver struct for resolving services
type serviceResolver struct {
	service *models.Service
}

// ID service id
func (r *serviceResolver) ID() graphql.ID {
	ID := strconv.Itoa(int(r.service.ID))

	return graphql.ID(ID)
}

// Name service name
func (r *serviceResolver) Name() *string {
	return &r.service.Name
}

// Duration service duration
func (r *serviceResolver) Duration() *int32 {
	return &r.service.Duration
}

// Price service name
func (r *serviceResolver) Price() *int32 {
	return &r.service.Price
}

// Trend service trend
func (r *serviceResolver) Trend() *string {
	return &r.service.Trend
}

// ProviderID service name
func (r *serviceResolver) ProviderID() *graphql.ID {
	ID := strconv.Itoa(int(r.service.ProviderID))

	providerID := graphql.ID(ID)
	return &providerID
}

// Status service name
func (r *serviceResolver) Status() *bool {
	return &r.service.Status
}

// Services base on user query
func (r *Resolver) Services(ctx context.Context, args getServiceQuery) (*[]*serviceResolver, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)

	// set default limit and offset
	var limit int32
	var offset int32 = 1

	if args.Limit == nil || (args.Limit != nil && *args.Limit == 0) {
		limit = config.ServicesLimit
	} else {
		limit = *args.Limit
	}

	if args.Page != nil && *args.Page > 0 {
		offset = *args.Page
	}

	offset = limit * (offset - 1)

	services, err := svc.Datastore.ServiceDB.GetServices(
		args.Name,
		*args.Latitude,
		*args.Longitude,
		*args.Radius,
	)
	if err != nil {
		return nil, err
	}

	var results []*serviceResolver

	for _, service := range services {
		results = append(results, &serviceResolver{service})
	}
	return &results, nil
}
