package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID string `gorm:"primary_key"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartDBService struct {
	DB *gorm.DB
}
