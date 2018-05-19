package resolver

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/models"
)

// userResolver struct for resolving users
type userResolver struct {
	u *models.User
}

// ID user id
func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

// Name user name
func (r *userResolver) Name() *string {
	return &r.u.Name
}

// User user
func (r *Resolver) User(ctx context.Context) *userResolver {
	return &userResolver{&models.User{
		ID:   "1",
		Name: "Enaho Murphy",
	}}
}
