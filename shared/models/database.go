package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gitlab.com/synergy-designs/style-blitz/shared/config"
)

// DB database abstraction
type DB struct {
	*gorm.DB
}

var database *gorm.DB

// Connect connects to the database connection
func Connect(conf *config.Config) *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
	db, err := gorm.Open("mysql", dbURL)
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

	return &Datastore {
		ProviderDB: &ProviderDbService{DB},
		UserDB: &UserDbService{DB},
		ServiceDB: &ServiceDBService{DB},
	}
}

type Datastore struct {
	ProviderDB ProviderDB
	UserDB UserDB
	ServiceDB ServiceDB
}