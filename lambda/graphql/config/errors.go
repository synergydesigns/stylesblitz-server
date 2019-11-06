package config

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func AuthenticationError(cxt context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Message: "you need to be authenticated to access this data",
		Extensions: map[string]interface{}{
			"code":   401,
			"status": "Authentication",
		},
		Path: graphql.GetResolverContext(cxt).Path(),
	}
}
