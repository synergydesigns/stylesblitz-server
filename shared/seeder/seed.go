package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/jinzhu/gorm"
	"gitlab.com/synergy-designs/style-blitz/shared/config"
	"gitlab.com/synergy-designs/style-blitz/shared/models"
)

var conf = config.LoadConfig()

// Seeder Struct that handles databse seeding
type Seeder struct {
	Table  string
	DB     *gorm.DB
	File   []byte
	ROOT   string
	Tables []string
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
	case "category":
		var data []models.Category

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Category) {
				fmt.Println(v)
				s.DB.Table(schema).Create(&v)
			}(v)
		}
		break
	case "provider":
		var data []models.Provider

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Provider) {
				fmt.Println(v)
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
	case "service":
		var data []models.Service

		json.Unmarshal(s.File, &data)
		for _, v := range data {
			func(v models.Service) {
				fmt.Println(v)
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
			fmt.Println(value, "==========")
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

// func main() {
// 	seed := Seeder{}

// 	seed.SetTables([]string{"provider", "address", "category", "service"})

// 	seed.Init().Clean()
// 	// seed.LoadData("provider").Seed("provider")
// 	// seed.LoadData("address").Seed("address")
// 	// seed.LoadData("category").Seed("category")
// 	seed.LoadData("service").Seed("service")
// }