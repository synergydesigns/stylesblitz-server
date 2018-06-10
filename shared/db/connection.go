package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	dbName     = os.Getenv("MYSQL_DATABASE")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbUser     = os.Getenv("MYSQL_USER")
	dbHost     = os.Getenv("DATABASE_HOST")
)

// Connection holds the database connection
func Connection() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
	fmt.Println(dbURL)
	connection, err := gorm.Open("mysql", dbURL)

	// disable database pluralization
	connection.SingularTable(true)

	if err != nil {
		log.Fatal(err)
	}

	return connection
}
