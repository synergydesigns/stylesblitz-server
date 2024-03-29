package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
)

type DB struct {
	*gorm.DB
}

var database *gorm.DB

func Connect(conf *config.Config) *gorm.DB {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", conf.DBUser,
		conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName, conf.SSLMode,
	)
	db, err := gorm.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	database = db

	if conf.GoEnv == "development" {
		db.LogMode(true)
	}

	return database
}

func NewDB(config *config.Config) *Datastore {
	DB := Connect(config)

	return &Datastore{
		VendorDB:         &VendorDbService{DB},
		UserDB:           &UserDbService{DB},
		ServiceDB:        &ServiceDBService{DB},
		AssetDB:          &AssetDBService{DB},
		VendorCategoryDB: &VendorCategoryDBService{DB},
		CartDB:           &CartDBService{DB},
		ProductDB:        &ProductDBService{DB},
		AutocompleteDB:   &AutocompleteDBService{DB},
		ServiceReviewDB:  &ServiceReviewDBService{DB},
	}
}

type Datastore struct {
	VendorDB         VendorDB
	UserDB           UserDB
	AssetDB          AssetDB
	ServiceDB        ServiceDB
	VendorCategoryDB VendorCategoryDB
	CartDB           CartDB
	ProductDB        ProductDB
	AutocompleteDB   AutocompleteDB
	ServiceReviewDB  ServiceReviewDB
}
