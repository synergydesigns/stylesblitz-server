package models

import graphql "github.com/graph-gophers/graphql-go"

// Service defines the service models for graphql
// for getting a single service
type Service struct {
	ID       graphql.ID
	Name     string
	Duration string
	Price    string
	Status   string
	Trend    string
	ShopID   graphql.ID
}
