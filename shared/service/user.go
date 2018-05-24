package service

import (
	"errors"
	"fmt"
	"log"

	"gitlab.com/synergy-designs/style-blitz/shared/db"
)

func GetUserByID(id uint) (db.User, error) {
	var user db.User
	result := DB.First(&user, id)

	if result.Error != nil {
		log.Printf("Could not find User: %v", result.Error)
		return user, errors.New(fmt.Sprintf("User with id %d cannot be found", id))
	}

	defer DB.Close()

	return user, nil
}
