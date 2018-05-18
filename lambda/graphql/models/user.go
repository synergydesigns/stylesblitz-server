package models

import graphql "github.com/graph-gophers/graphql-go"

type User struct {
	ID   graphql.ID
	Name string
}
