package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/synergydesigns/stylesblitz-server/shared/seeder"

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

	if err != nil {
		log.Fatal("Error connecting to db", err)
	}

	driver, _ := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		migrationFolder,
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("Error create new migration instance", err)
	}

	// start migration
	if os.Args[1] == "up" {
		err = m.Up()

		if err != nil {
			log.Fatal("Error running migration", err)
		}
	}

	// tear down migration
	if os.Args[1] == "down" {
		err = m.Down()

		if err != nil {
			log.Fatal("Error running migration", err)
		}
	}

	// tear down migration
	if os.Args[1] == "drop" {
		err = m.Drop()

		if err != nil {
			log.Fatal("Error droping migrations", err)
		}
	}

	if os.Args[1] == "seed" {
		if os.Args[2] == "countries" {
			seeder.SeedCountries()
		}

		if os.Args[2] == "states" {
			seeder.SeedStates()
		}

		if os.Args[2] == "cities" {
			seeder.SeedCities()
		}

		if os.Args[2] == "locations" {
			seeder.SeedLocations()
		}

		if os.Args[2] == "vendors_all" {
			seeder.SeedVendorData()
		}

		if os.Args[2] == "services" {
			seeder.New().LoadData("services").Seed("services")
		}

		if os.Args[2] == "categories" {
			seeder.New().LoadData("categories").Seed("categories")
		}

		if os.Args[2] == "vendors" {
			seeder.New().LoadData("vendors").Seed("vendors")
		}

		if os.Args[2] == "address" {
			seeder.New().LoadData("address").Seed("address")
		}
	}
}
