package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connection holds the database connection
func Connection() *gorm.DB {
	dbURL := fmt.Sprintf("%s:@/styleblitz", "root")
	connection, err := gorm.Open("mysql", dbURL)

	// disable database pluralization
	connection.SingularTable(true)

	if err != nil {
		log.Fatal(err, connection)
	}

	return connection
}
