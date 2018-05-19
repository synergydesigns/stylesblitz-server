package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/synergy-designs/style-blitz/migrations/utils"
)

func main() {
	config := utils.LoadConfig()
	migrations := utils.GenarateDarwinMigrations(config)

	fmt.Println(config, migrations)
	dbURL := fmt.Sprintf("%s:@/styleblitz", config.DBUser)
	database, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	driver := darwin.NewGenericDriver(database, darwin.MySQLDialect{})
	d := darwin.New(driver, migrations, nil)
	err = d.Migrate()

	if err != nil {
		log.Println(err)
	}
}
