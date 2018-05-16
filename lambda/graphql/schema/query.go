package schema

import (
	"github.com/graphql-go/graphql"
)

// QueryType defines graphql queries
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return User{ID: 3, Name: "Enaho Murphy"}, nil
				},
			},
		},
	},
)
