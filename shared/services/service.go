package service

import (
	config "github.com/synergydesigns/stylesblitz-server/shared/config"
	models "github.com/synergydesigns/stylesblitz-server/shared/models"
)

// Services Holds all methods that futher abstract
// database integration
type Services struct {
	Datastore *models.Datastore
}

// New initializes all services
func New() *Services {
	conf := config.LoadConfig()
	return &Services{
		Datastore: models.NewDB(conf),
	}
}
