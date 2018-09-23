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
func Connect(conf *config.Config) *DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	// we want to cache db connection
	database = db

	// disable database pluralization
	database.SingularTable(true)

	return &DB{database}
}

// NewDB initializes the database instance
func NewDB(config *config.Config) *DB {
	return Connect(config)
}

// Datastore defines all the methods used to
// interface with the database
type Datastore interface {
	GetUserByID(id uint64) (*User, error)
}
