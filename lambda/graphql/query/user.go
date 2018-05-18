package query

import (
	"context"

	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/models"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/resolver"
)

func (r *resolver.Resolver) User(ctx context.Context) *resolver.UserResolver {
	return &resolver.UserResolver{models.User{
		ID:   "1",
		Name: "Enaho Murphy",
	}}
}
