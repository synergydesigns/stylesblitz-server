package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/models"
)

type User struct {
	ID   graphql.ID
	Name string
}

// UserResolver struct for resolving users
type UserResolver struct {
	u *models.User
}

func (r *UserResolver) ID() graphql.ID {
	return r.u.ID
}
