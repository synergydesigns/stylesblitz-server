package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/migrations/utils"
)

// Migrate migrates db
func Migrate() {
	config := config.LoadConfig()
	migrations := utils.GenarateDarwinMigrations(config)

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser,
		config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	database, err := sql.Open("mysql", dbURL+"?parseTime=true")

	defer database.Close()

	if err != nil {
		log.Fatal(err)
	}

	driver := darwin.NewGenericDriver(database, darwin.MySQLDialect{})
	d := darwin.New(driver, migrations, nil)
	err = d.Migrate()

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Migration completed")
	}
}

func main() {
	Migrate()
}
