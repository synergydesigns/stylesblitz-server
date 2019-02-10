package models

import (
	"fmt"
	"log"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
)

// DB database abstraction
type DB struct {
	*gorm.DB
}

var database *gorm.DB

// Connect connects to the database connection
func Connect(conf *config.Config) *gorm.DB {
	dbURL := fmt.Sprintf(
		"host=%d port=%d user=gorm dbname=%d password=%d",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBName, conf.DBPassword,
	)

	// db, err := gorm.Open("mysql", dbURL+"?parseTime=true")
	db, err := gorm.Open("postgres", dbURL)
	log.Println(dbURL)

	if err != nil {
		log.Fatal(err)
	}

	// we want to cache db connection
	database = db

	// disable database pluralization
	database.SingularTable(true)

	return database
}

// NewDB initializes the database instance
func NewDB(config *config.Config) *Datastore {
	DB := Connect(config)

	return &Datastore{
		ProviderDB: &ProviderDbService{DB},
		UserDB:     &UserDbService{DB},
		ServiceDB:  &ServiceDBService{DB},
	}
}

type Datastore struct {
	ProviderDB ProviderDB
	UserDB     UserDB
	ServiceDB  ServiceDB
}
