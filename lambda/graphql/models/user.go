package models

import graphql "github.com/graph-gophers/graphql-go"

// User defines the user models for graphql
// for getting a single user
type User struct {
	ID           graphql.ID
	Username     string
	Email        string
	Name         string
	Phone        int32
	Password     string
	ProfileImage string
	WallImage    string
	Bio          string
}
