package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/jinzhu/gorm"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

var conf = config.LoadConfig()

// Seeder Struct that handles database seeding
type Seeder struct {
	Table  string
	DB     *gorm.DB
	File   []byte
	ROOT   string
	Tables []string
}

// New initializes seeder with db configuration
func New() *Seeder {
	return new(Seeder).Init()
}

// LoadData Loads the data to be seeded
func (s *Seeder) LoadData(file string) *Seeder {
	s.File = getData(file)

	return s
}

// Init initialisesa DB connection
func (s *Seeder) Init() *Seeder {

	s.DB = models.Connect(conf)
	s.ROOT = conf.RootDirectory

	return s
}

// Seed seed the Loaded data
func (s *Seeder) Seed(schema string) *Seeder {
	switch schema {
	case "categories":
		var data []models.Category

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Category) {
				s.DB.Table(schema).Create(&v)
			}(v)
		}
		break
	case "vendors":
		var data []models.Vendor

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Vendor) {
				s.DB.Table(schema).Create(&v)
			}(v)
		}
		break
	case "address":
		var data []models.Address

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Address) {
				fmt.Println(v)
				s.DB.Table(schema).Create(&v)
			}(v)
		}
	case "services":
		var data []models.Service

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Service) {
				s.DB.Table(schema).Create(&v)
			}(v)
		}
		break
	case "assets":
		var data []models.Asset

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Asset) {
				s.DB.Table(schema).Create(&v)
			}(v)
		}
	}

	s.Tables = append(s.Tables, schema)
	return s
}

func (s *Seeder) Clean() *Seeder {
	for _, value := range s.Tables {

		if s.DB.HasTable(value) {
			s.DB.Exec("TRUNCATE TABLE " + value)
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
	fileBytes, err := ioutil.ReadFile(path.Join(conf.RootDirectory, "shared/seeder/"+file+".json"))

	if err != nil {
		log.Fatal(err)
	}

	return fileBytes
}

func (s *Seeder) SetTables(tables []string) *Seeder {
	s.Tables = tables

	return s
}

func SeedVendorData() {
	seed := New()
	schemas := []string{
		"address",
		// "assets",
		// "users",
		// "vendors",
	}

	for _, schema := range schemas {
		seed.LoadData(schema)
		seed.Seed(schema)
	}
}
