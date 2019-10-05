package models

import (
	"time"

	"github.com/lucsky/cuid"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID string `gorm:"primary_key"`
	UserID string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartDBService struct {
	DB *gorm.DB
}

func (cart *Cart) BeforeCreate(scope *gorm.Scope) error {
	if cart.ID == "" {
		scope.SetColumn("ID", cuid.New())
	}

	return nil
}

type CartDB interface {
	Cart(userID string) (string, error)
}
