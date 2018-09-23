package resolver

import (
	"context"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/config"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
	service "gitlab.com/synergy-designs/style-blitz/shared/services"
)

// userResolver struct for resolving users
type userResolver struct {
	u *models.User
}

// ID user id
func (r *userResolver) ID() graphql.ID {
	ID := strconv.Itoa(int(r.u.ID))

	return graphql.ID(ID)
}

// Name user name
func (r *userResolver) Name() *string {
	return &r.u.Name
}

// Email user name
func (r *userResolver) Email() *string {
	return &r.u.Email
}

// Email user name
func (r *userResolver) Phone() *int32 {
	return &r.u.Phone
}

// Email user name
func (r *userResolver) Password() *string {
	return &r.u.Email
}

// Email user name
func (r *userResolver) ProfileImage() *string {
	return &r.u.ProfileImage
}

// Email user name
func (r *userResolver) WallImage() *string {
	return &r.u.WallImage
}

// Email user name
func (r *userResolver) Bio() *string {
	return &r.u.Bio
}

// User user
func (r *Resolver) User(ctx context.Context, args struct {
	ID   string
	Name string
}) (*userResolver, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	userID, _ := strconv.ParseUint(args.ID, 10, 64)
	user, err := svc.Datastore.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &userResolver{user}, nil
}
