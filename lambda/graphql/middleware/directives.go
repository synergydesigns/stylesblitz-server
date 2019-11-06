package middleware

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
)

func IsAuthenticated(
	ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user := config.GetUser(ctx)

	if user == nil {
		return false, config.AuthenticationError(ctx)
	}

	return next(ctx)
}
