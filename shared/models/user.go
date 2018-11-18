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
	ID           uint64 `gorm:"primary_key"`
	Username     string
	Email        string
	Name         string
	Phone        int32
	Password     string
	ProfileImage string
	WallImage    string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserDbService struct {
	DB *gorm.DB
}

type UserDB interface {
	GetUserByID(id uint64) (*User, error)
}

// GetUserByID gets a single user by ID
// @params {userID} userID is an integer
func (service *UserDbService) GetUserByID(id uint64) (*User, error) {
	var user User

	result := service.DB.First(&user, id)

	if result.Error != nil {
		log.Printf("Could not find User: %v", result.Error)
		return nil, fmt.Errorf("User with id %d cannot be found", id)
	}

	return &user, nil
}
