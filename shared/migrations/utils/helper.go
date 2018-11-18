package utils

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/GuiaBolso/darwin"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	yaml "gopkg.in/yaml.v2"
)

// Migration is a struct for parsing sql files
type Migration struct {
	ChangeSet []ChangeSet `yaml:"changeSet"`
}

// ChangeSet is a struct defines the user that created the sql statement
// And the sql statement to eecute
type ChangeSet struct {
	ID          string  `yaml:"id"`
	Author      string  `yaml:"author"`
	Description string  `yaml:"description"`
	Changes     []Query `yaml:"changes"`
}

// Query defines the sql query to run
type Query struct {
	SQL string `yaml:"sql"`
}

// GenarateDarwinMigrations reads the json files in the query folder
// and generates migrations out of the json files found
func GenarateDarwinMigrations(config *config.Config) []darwin.Migration {
	var migrations []darwin.Migration

	filePath := path.Join(config.RootDirectory, config.MigrationPath)
	files, err := ioutil.ReadDir(filePath)

	if err != nil {
		panic(err)
	}

	version := 1
	for _, file := range files {
		migration := Migration{}

		fileBytes, _ := ioutil.ReadFile(path.Join(filePath, file.Name()))

		fmt.Println(string(fileBytes))

		if err := yaml.Unmarshal(fileBytes, &migration); err != nil {

			panic(err)
		}

		for _, changeSet := range migration.ChangeSet {
			for _, change := range changeSet.Changes {
				t := darwin.Migration{
					Version:     float64(version),
					Description: changeSet.Description,
					Script:      change.SQL,
				}
				version++
				migrations = append(migrations, t)
			}
		}
	}
	return migrations
}
