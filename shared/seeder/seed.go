package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gitlab.com/synergy-designs/style-blitz/shared/config"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
)

// Seeder Struct that handles databse seeding
type Seeder struct {
	Table string
	DB    *models.DB
	File  []byte
}

// LoadData Loads the data to be seeded
func (s *Seeder) LoadData(file string) *Seeder {
	s.File = getData(file)

	return s
}

// Init initialisesa DB connection
func (s *Seeder) Init() *Seeder {
	conf := config.LoadConfig()

	s.DB = models.NewDB(conf)

	return s
}

// Seed seed the Loaded data
func (s *Seeder) Seed(schema string) *Seeder {
	switch schema {
	case "category":
		var data []models.Category

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Category) {
				fmt.Println(v)
				s.DB.Table(schema).Create(&v)
			}(v)
		}
	}

	return s
}

// Close closes the database connection
func (s *Seeder) Close() *Seeder {
	s.DB.Close()
	return s
}

func getData(file string) []byte {
	dir, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(path.Join(dir, "shared/seeder/"+file+".json"))

	if err != nil {
		log.Fatal(err)
	}

	return fileBytes
}
