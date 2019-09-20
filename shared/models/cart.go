package model

import (
	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID string `gorm:"primary_key"`
	UserID int `json:"user_id"`
	CreatedAt bool `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type CartDBService struct {
	DB *gorm.DB
}
