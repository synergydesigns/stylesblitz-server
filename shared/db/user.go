package db

import (
	"time"
)

// User defines the user models for graphql
// for getting a single user
type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	Name         string
	Phone        string
	Password     string
	ProfileImage string
	WallImage    string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
