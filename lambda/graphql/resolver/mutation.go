package resolver

import (
	"context"
	"fmt"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

type mutationResolver struct{}

func (mutation mutationResolver) Login(ctx context.Context, email string, password string) (*string, error) {
	svc := ctx.Value(config.CTXKeyservices).(*service.Services)
	token := ""
	user, err := svc.Datastore.UserDB.GetUserByEmail(email)

	if err != nil {
		return &token, err
	}

	if !svc.JWT.CheckPasswordHash(password, user.Password) {
		return &token, fmt.Errorf("invalid email or password")
	}

	token, err = svc.JWT.GenerateAuthToken(*user)

	if err != nil {
		return &token, err
	}

	return &token, nil
}
