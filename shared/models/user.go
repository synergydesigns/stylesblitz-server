package models

import (
	"fmt"
	"log"
	"time"
)

// User defines the user models for graphql
// for getting a single user
type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	Name         string
	Phone        int
	Password     string
	ProfileImage string
	WallImage    string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// GetUserByID gets a single user by ID
// @params {userID} userID is an integer
func (db *DB) GetUserByID(id uint) (User, error) {
	var user User

	result := db.First(&user, id)

	if result.Error != nil {
		log.Printf("Could not find User: %v", result.Error)
		return user, fmt.Errorf("User with id %d cannot be found", id)
	}

	return user, nil
}
