package models_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
)

type AutocompleteServiceTestSuite struct {
	suite.Suite
	seed    *seeder.Seeder
	service models.AutocompleteDBService
}

func (suite *AutocompleteServiceTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.service = models.AutocompleteDBService{models.Connect(config.LoadConfig())}
}

func (suite *AutocompleteServiceTestSuite) TearDownTest() {
	suite.seed.Clean()
}

func (suite *AutocompleteServiceTestSuite) AfterTest() {
	suite.seed.Truncate("services").Truncate("categories")
}

func (suite *AutocompleteServiceTestSuite) TestGetSuggestions() {
	suite.seed.
		LoadAndSeed("users").
		LoadAndSeed("vendors").
		LoadAndSeed("address").
		LoadAndSeed("categories").
		LoadAndSeed("services")

	testCases := []struct {
		title          string
		query          string
		expectedLength int
		serviceCount   int
		vendorCount    int
		categoryCount  int
	}{
		{
			title:          "Should return 11 results matching relax",
			query:          "relax",
			expectedLength: 11,
			serviceCount:   9,
			vendorCount:    1,
			categoryCount:  1,
		},
		{
			title:          "Should return 12 result matching beauty",
			query:          "beauty",
			expectedLength: 3,
			serviceCount:   0,
			vendorCount:    3,
			categoryCount:  0},
		{
			title:          "Should return 6 result matching saloon",
			query:          "saloon",
			expectedLength: 6,
			serviceCount:   0,
			vendorCount:    2,
			categoryCount:  4,
		},
	}

	for _, test := range testCases {
		suggestions, err := suite.service.GetSuggestions(test.query)

		var serviceCount int
		var vendorCount int
		var categoryCount int

		for _, suggestion := range suggestions {
			if suggestion.Type == "vendors" {
				vendorCount += 1
			}

			if suggestion.Type == "services" {
				serviceCount += 1
			}

			if suggestion.Type == "categories" {
				categoryCount += 1
			}
		}

		suite.Nil(err, test.title)
		suite.Equal(len(suggestions), test.expectedLength, test.title)
		suite.Equal(serviceCount, test.serviceCount, test.title)
		suite.Equal(vendorCount, test.vendorCount, test.title)
		suite.Equal(categoryCount, test.categoryCount, test.title)
	}
}

func TestAutocompleteServiceSuite(t *testing.T) {
	suite.Run(t, new(AutocompleteServiceTestSuite))
}
