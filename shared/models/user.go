package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// User defines the user models for graphql
// for getting a single user
type User struct {
	ID           string `gorm:"primary_key"`
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	Password     string
	Bio          string
	Phone        string
	ProfileImage string
	WallImage    string
	AddressID    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserDbService struct {
	DB *gorm.DB
}

type UserDB interface {
	GetUserByID(id string) (*User, error)
}

// GetUserByID gets a single user by ID
// @params {userID} userID is an integer
func (service *UserDbService) GetUserByID(id string) (*User, error) {
	var user User

	result := service.DB.First(&user, id)

	if result.Error != nil {
		log.Printf("Could not find User: %v", result.Error)
		return nil, fmt.Errorf("User with id %s cannot be found", id)
	}

	return &user, nil
}
