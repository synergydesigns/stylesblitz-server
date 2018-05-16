package schema

import "github.com/graphql-go/graphql"

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserType Defines User schema
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
