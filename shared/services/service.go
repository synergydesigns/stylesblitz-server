package service

import (
	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	models "github.com/synergydesigns/stylesblitz-server/shared/models"
)

type Services struct {
	Datastore *models.Datastore
	JWT       JWT
}

// New initializes all services
func New() *Services {
	conf := config.LoadConfig()
	return &Services{
		Datastore: models.NewDB(conf),
		JWT:       NewJWT(conf),
	}
}
