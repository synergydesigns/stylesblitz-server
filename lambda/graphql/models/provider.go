package models

import graphql "github.com/graph-gophers/graphql-go"

// Provider defines the provider models for graphql
// for getting a single provider
type Provider struct {
	ID          graphql.ID
	Name        string
	Description string
	About       string
	Phone       string
	UserID      graphql.ID
	AddressID   graphql.ID
}
