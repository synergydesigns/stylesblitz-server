package service

import (
	config "gitlab.com/synergy-designs/style-blitz/shared/config"
	models "gitlab.com/synergy-designs/style-blitz/shared/models"
)

// Services Holds all methods that futher abstract
// database integration
type Services struct {
	datastore models.Datastore
}

// New initializes all services
func New() *Services {
	conf := config.LoadConfig()
	return &Services{
		datastore: models.NewDB(conf),
	}
}
