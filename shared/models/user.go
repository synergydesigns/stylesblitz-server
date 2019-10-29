package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `gorm:"primary_key"`
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	Password     string `json:"password,omitempty"`
	Bio          string
	Phone        *string
	ProfileImage string
	WallImage    string
	AddressID    int
	Assets       []*Asset `gorm:"many2many:user_assets;"`
	Vendor       *Vendor  `gorm:"foreignkey:user_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserDbService struct {
	DB *gorm.DB
}

type UserDB interface {
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return err
	}

	err = scope.SetColumn("Firstname", strings.ToLower(user.Firstname))

	if err != nil {
		return err
	}

	err = scope.SetColumn("Lastname", strings.ToLower(user.Lastname))

	if err != nil {
		return err
	}

	err = scope.SetColumn("Email", strings.ToLower(user.Email))

	if err != nil {
		return err
	}

	err = scope.SetColumn("Password", password)

	if user.ID == "" {
		err = scope.SetColumn("ID", cuid.New())
	}

	return err
}

func (service *UserDbService) GetUserByID(id string) (*User, error) {
	var user User
	result := service.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Printf("Could not find User: %v", result.Error)
		return nil, fmt.Errorf("user with id %s cannot be found", id)
	}

	return &user, nil
}

func (service *UserDbService) GetUserByEmail(email string) (*User, error) {
	var user User
	result := service.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Printf("Could not find User: %v", result.Error)
		return nil, fmt.Errorf("user with email %s cannot be found", email)
	}

	return &user, nil
}
