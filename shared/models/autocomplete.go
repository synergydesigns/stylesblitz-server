package models

import "github.com/jinzhu/gorm"

type Autocomplete struct {
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

func (service *AutocompleteDBService) GetSuggestions(query string) ([]*Autocomplete, error) {
	var autocomplete []*Autocomplete

	service.DB.Where("tsv @@ plainto_tsquery(?)", query).Find(&autocomplete)

	return autocomplete, nil
}
