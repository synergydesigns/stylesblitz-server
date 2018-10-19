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
type ServiceResolver struct {
	service *models.Service
}

// ID service id
func (r *ServiceResolver) ID() graphql.ID {
	ID := strconv.Itoa(int(r.service.ID))

	return graphql.ID(ID)
}

// Name service name
func (r *ServiceResolver) Name() *string {
	return &r.service.Name
}

// Duration service duration
func (r *ServiceResolver) Duration() *int32 {
	return &r.service.Duration
}

// Price service name
func (r *ServiceResolver) Price() *int32 {
	return &r.service.Price
}

// Trend service trend
func (r *ServiceResolver) Trend() *string {
	return &r.service.Trend
}

// ProviderID service name
func (r *ServiceResolver) ProviderID() *graphql.ID {
	ID := strconv.Itoa(int(r.service.ProviderID))

	providerID := graphql.ID(ID)
	return &providerID
}

// Status service name
func (r *ServiceResolver) Status() *bool {
	return &r.service.Status
}

type getServiceQuery struct {
	Longitude, Latitude, Radius *float64
	Name                        string
}

// Services base on user query
func (r *Resolver) Services(ctx context.Context, args getServiceQuery) (*[]*ServiceResolver, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	services, err := svc.Datastore.GetServices(
		args.Name,
		*args.Latitude,
		*args.Longitude,
		*args.Radius,
	)
	if err != nil {
		return nil, err
	}

	var results []*ServiceResolver

	for _, service := range services {
		results = append(results, &ServiceResolver{service})
	}
	return &results, nil
}
