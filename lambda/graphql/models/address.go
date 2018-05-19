package models

import graphql "github.com/graph-gophers/graphql-go"

// Address defines the address models for graphql
// for getting a single address
type Address struct {
	ID      graphql.ID
	ShopID  graphql.ID
	Country string
	State   string
	City    string
	ZipCode string
}
