package resolver

import (
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/genql"
)

type Resolver struct{}

func (r *Resolver) Vendor() genql.VendorResolver {
	return &VendorResolver{r}
}

func (r *Resolver) Query() genql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Address() genql.AddressResolver {
	return &addressResolver{r}
}

func (r *Resolver) Service() genql.ServiceResolver {
	return &serviceResolver{}
}

func (r *Resolver) Mutation() genql.MutationResolver {
	return &mutationResolver{}
}
