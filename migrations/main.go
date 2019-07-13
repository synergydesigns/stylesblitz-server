package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
)

const migrationFolder = "file://migrations/query"

func main() {
	config := config.LoadConfig()

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", config.DBUser,
		config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.SSLMode,
	)
	db, err := sql.Open("postgres", dbURL)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	driver, _ := postgres.WithInstance(db, &postgres.Config{})

	// start migration
	if os.Args[1] == "up" {
		fmt.Println("Hello world")
		m, err := migrate.NewWithDatabaseInstance(
			migrationFolder,
			"postgres", driver)
		m.Steps(100)

		if err != nil {
			log.Fatal(err)
		}
	}

	// tear down migration
	if os.Args[1] == "down" {
		m, err := migrate.New(migrationFolder, dbURL)

		err = m.Down()

		if err != nil {
			log.Fatal(err)
		}
	}
}
