package resolver

import (
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/genql"
)

type Resolver struct{}

func (r *Resolver) Provider() genql.ProviderResolver {
	return &providerResolver{r}
}
func (r *Resolver) Query() genql.QueryResolver {
	return &queryResolver{r}
}
