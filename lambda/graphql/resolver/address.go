package resolver

import (
	"context"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
)

type addressResolver struct {
	address *models.Address
}

// ID is a unique identifier for address
func (r *addressResolver) ID() graphql.ID {
	ID := strconv.Itoa(int(r.address.ID))

	return graphql.ID(ID)
}

// ProviderID signifies the provider that owns this address
func (r *addressResolver) ProviderID() graphql.ID {
	ID := strconv.Itoa(int(r.address.ProviderID))

	return graphql.ID(ID)
}

// Country the provider shop/service center can be found
func (r *addressResolver) Country() *string {
	return &r.address.Country
}

// state the provider shop/service center can be found
func (r *addressResolver) State() *string {
	return &r.address.State
}

// City the provider shop/service center can be found
func (r *addressResolver) City() *string {
	return &r.address.City
}

// City the provider shop/service center can be found
func (r *addressResolver) ZipCode() *string {
	return &r.address.ZipCode
}

// Latitude is the first half of the geo point of this address
func (r *addressResolver) Latitude() *float64 {
	return &r.address.Latitude
}

// Longitude is the second half of the geo point of this address
func (r *addressResolver) Longitute() *float64 {
	return &r.address.Longitute
}

func (r *Resolver) Addresses(ctx context.Context, args struct {
	ProviderID uint
}) *[]*addressResolver {
	var addresses []*addressResolver

	addresses = append(addresses, &addressResolver{&models.Address{}})
	return &addresses
}

func (r *Resolver) GetAddress(ctx context.Context, args struct {
	ID int32
}) *addressResolver {
	return &addressResolver{&models.Address{}}
}
