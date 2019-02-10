package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/synergydesigns/stylesblitz-server/migrations/utils"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
)

// Migrate migrates db
func Migrate() {
	config := config.LoadConfig()
	migrations := utils.GenarateDarwinMigrations(config)

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", config.DBUser,
		config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.SSLMode,
	)

	database, err := sql.Open("postgres", dbURL)

	defer database.Close()

	if err != nil {
		log.Fatal(err)
	}

	driver := darwin.NewGenericDriver(database, darwin.PostgresDialect{})
	d := darwin.New(driver, migrations, nil)

	infos, err := d.Info()
	if err != nil {
		fmt.Println(err)
	}

	for _, info := range infos {
		fmt.Printf("%.1f %s %s\n", info.Migration.Version, info.Status, info.Migration.Description)
	}

	err = d.Migrate()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Migration completed")
	}
}

func main() {
	Migrate()
}
