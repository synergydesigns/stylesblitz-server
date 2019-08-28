package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Asset struct {
	ID          string `gorm:"primary_key"`
	Title       string
	Description string
	Caption     string
	Alt         string
	MediaType   string
	MimeType    string
	Width       int
	Height      int
	Filename    string
	Size        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AssetDBService struct {
	DB *gorm.DB
}

type AssetDB interface {
	CreateAsset(asset Asset, id string) (*Asset, error)
	GetAsset(id string) (Asset, error)
}

func (service *AssetDBService) CreateAsset(payload Asset, id string) (*Asset, error) {
	panic(1)
}

func (service *AssetDBService) GetAsset(id string) (Asset, error) {
	panic(1)
}
