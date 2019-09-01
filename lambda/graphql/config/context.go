package config

import (
	"context"

	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

type contextKey string

const CTXKeyservices = contextKey("services")
const CTXKeyuser = contextKey("user")

func GetServices(ctx context.Context) *service.Services {
	return ctx.Value(CTXKeyservices).(*service.Services)
}
