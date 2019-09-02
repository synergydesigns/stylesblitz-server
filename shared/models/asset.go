package models

import (
	"fmt"
	"log"
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
	Key         string
	Bucket      string
	Url         string
	User        []User   `gorm:"many2many:user_assets;association_autoupdate:false;association_autocreate:false"`
	Vendor      []Vendor `gorm:"many2many:user_assets;association_autoupdate:false;association_autocreate:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AssetDBService struct {
	DB *gorm.DB
}

type AssetDB interface {
	CreateAsset(asset Asset, id string) (*Asset, error)
	GetAsset(id string) (Asset, error)
	CreateAssets(assets []Asset) ([]Asset, error)
}

func (service *AssetDBService) CreateAsset(payload Asset, id string) (*Asset, error) {
	panic(1)
}

func (service *AssetDBService) GetAsset(id string) (Asset, error) {
	panic(1)
}

func (service *AssetDBService) CreateAssets(assets []Asset) ([]Asset, error) {
	for _, asset := range assets {
		result := service.DB.Create(&asset)

		if result.Error != nil {
			log.Printf("Could not find User: %v", result.Error)
			return nil, fmt.Errorf("Error occurred creating assets")
		}
	}

	return assets, nil
}
