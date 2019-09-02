package service

import (
	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	models "github.com/synergydesigns/stylesblitz-server/shared/models"
)

type Services struct {
	Datastore *models.Datastore
	AWS       AWS
	JWT       JWT
}

// New initializes all services
func New(conf *config.Config) *Services {
	return &Services{
		Datastore: models.NewDB(conf),
		AWS:       NewAWS(conf),
		JWT:       NewJWT(conf),
	}
}
