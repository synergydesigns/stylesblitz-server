package service

import (
	"gitlab.com/synergy-designs/style-blitz/shared/db"
)

// Service Holds all methods that futher abstract
// database integration
type Service struct{}

// DB Holds all methods for interfacing
// with the database
var DB = db.Connection()
