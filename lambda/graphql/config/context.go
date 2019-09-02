package config

import (
	"context"

	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

type contextKey string

const CTXKeyservices = contextKey("services")
const CTXKeyuser = contextKey("user")
const CTConfig = contextKey("config")

func GetServices(ctx context.Context) *service.Services {
	return ctx.Value(CTXKeyservices).(*service.Services)
}

func GetUser(ctx context.Context) models.User {
	return ctx.Value(CTXKeyuser).(models.User)
}

func GetConfig(ctx context.Context) *config.Config {
	return ctx.Value(CTConfig).(*config.Config)
}
