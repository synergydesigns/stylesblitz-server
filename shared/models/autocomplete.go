package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Autocomplete struct {
	ID    string
	Query string
	Type  string
	Url   string
}

type AutocompleteDBService struct {
	DB *gorm.DB
}

type AutocompleteDB interface {
	GetSuggestions(query string) ([]*Autocomplete, error)
}

func (Autocomplete) TableName() string {
	return "autocomplete"
}

func (service *AutocompleteDBService) GetSuggestions(query string) ([]*Autocomplete, error) {
	var autocomplete []*Autocomplete

	result := service.DB.Where("tsv @@ plainto_tsquery(?)", query).Find(&autocomplete)

	if result.Error != nil {
		log.Printf("An error occurred getting suggestions %v", result.Error.Error())
		return autocomplete, fmt.Errorf("An error occurred getting suggestions %s", result.Error.Error())
	}

	return autocomplete, nil
}
