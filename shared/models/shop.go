package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lucsky/cuid"
)

type Shop struct {
	ID string `gorm:"primary_key"`
	Name string `json:"name"`
	VendorID string `json:"vendor_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShopDBService struct {
	DB *gorm.DB
}

func (shop *Shop) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", cuid.New())

	return err
}
