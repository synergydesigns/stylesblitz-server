package models_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

var categoryService = models.ServiceDBService{models.Connect(config.LoadConfig())}

type CategoryServiceTestSuite struct {
	suite.Suite
	vendor    models.Vendor
	vendor2   models.Vendor
	category  models.VendorCategory
	category2 models.VendorCategory
	category3 models.VendorCategory
	user      models.User
	seed      *seeder.Seeder
}

func (suite *CategoryServiceTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.seed.Tables = []string{"categories", "vendors", "users", "services"}
	suite.user = suite.seed.SeedUser("", "testuserservice", "testuserservice@gmail.com", nil)
	suite.vendor = suite.seed.SeedVendor("", suite.user.ID, "new custom vendor")
	suite.vendor2 = suite.seed.SeedVendor("", suite.user.ID, "new custom vendor 2")
	suite.vendor2 = suite.seed.SeedVendor("", suite.user.ID, "new custom vendor 3")
	suite.category = suite.seed.VendorCategory(1, suite.vendor.ID, "native")
	suite.category2 = suite.seed.VendorCategory(2, suite.vendor.ID, "make up")
	suite.category3 = suite.seed.VendorCategory(3, suite.vendor2.ID, "make up")
}

func (suite *CategoryServiceTestSuite) TearDownTest() {
	suite.seed.Clean()
}

func (suite *CategoryServiceTestSuite) AfterTest() {
	suite.seed.Truncate("services").Truncate("categories")
}

func (suite *CategoryServiceTestSuite) TestGetServicesByVendor() {
	for _, value := range utils.MakeRange(1, 30) {
		data := models.ServiceInput{
			Name:         fmt.Sprintf("braiding %d", value),
			Price:        utils.Float64ToPointer(40),
			Duration:     40,
			DurationType: "mins",
			Trending:     utils.BoolToPointer(true),
			CategoryID:   int(suite.category.ID),
		}

		if value <= 20 {
			data.VendorID = suite.vendor.ID
			suite.seed.VendorService(
				0,
				data,
			)
		} else {
			data.VendorID = suite.vendor2.ID
			suite.seed.VendorService(
				0,
				data,
			)
		}
	}
	testcases := []struct {
		title         string
		vendorID      string
		expectedCount int
	}{
		{
			title:         "Should get all 20 services for a vendor 1",
			vendorID:      suite.vendor.ID,
			expectedCount: 20,
		},
		{
			title:         "Should get all 10 services for a vendor 2",
			vendorID:      suite.vendor2.ID,
			expectedCount: 10,
		},
	}

	for _, test := range testcases {
		services, err := categoryService.GetServicesByVendor(test.vendorID)
		suite.Nil(err)
		suite.Equal(len(services), test.expectedCount, test.title)
	}
}

func (suite *CategoryServiceTestSuite) TestGetServicesByCategory() {
	for _, value := range utils.MakeRange(1, 30) {
		data := models.ServiceInput{
			Name:         fmt.Sprintf("braiding %d", value),
			Price:        utils.Float64ToPointer(40),
			Duration:     40,
			DurationType: "mins",
			Trending:     utils.BoolToPointer(true),
		}

		if value <= 20 {
			data.VendorID = suite.vendor.ID
			data.CategoryID = int(suite.category.ID)
			suite.seed.VendorService(
				0,
				data,
			)
		} else {
			data.VendorID = suite.vendor2.ID
			data.CategoryID = int(suite.category3.ID)
			suite.seed.VendorService(
				0,
				data,
			)
		}
	}
	testcases := []struct {
		title         string
		vendorID      string
		expectedCount int
		categoryID    uint64
	}{
		{
			title:         fmt.Sprintf("Should get all 20 services for a category %d", suite.category.ID),
			vendorID:      suite.vendor.ID,
			categoryID:    suite.category.ID,
			expectedCount: 20,
		},
		{
			title:         fmt.Sprintf("Should get all 10 services for a category %d", suite.category3.ID),
			vendorID:      suite.vendor2.ID,
			categoryID:    suite.category3.ID,
			expectedCount: 10,
		},
		{
			title:         "Should return 0 services if category  does not belong to vendor",
			vendorID:      suite.vendor2.ID,
			categoryID:    suite.category.ID,
			expectedCount: 0,
		},
	}

	for _, test := range testcases {
		services, err := categoryService.GetServicesByCategory(test.vendorID, test.categoryID)
		suite.Nil(err)
		suite.Equal(len(services), test.expectedCount, test.title)
	}
}

func (suite *CategoryServiceTestSuite) TestCreateService() {
	testCases := []struct {
		title         string
		payload       models.ServiceInput
		error         bool
		expectedError string
	}{
		{
			title: "Should create service with name braiding",
			payload: models.ServiceInput{
				Name:         fmt.Sprintf("braiding"),
				Price:        utils.Float64ToPointer(40),
				Duration:     40,
				DurationType: "mins",
				Trending:     utils.BoolToPointer(true),
				CategoryID:   int(suite.category.ID),
				VendorID:     suite.vendor.ID,
			},
		},
		{
			title: "Should create service with name make up for category2",
			payload: models.ServiceInput{
				Name:         fmt.Sprintf("make up"),
				Price:        utils.Float64ToPointer(40),
				Duration:     40,
				DurationType: "mins",
				Trending:     utils.BoolToPointer(true),
				CategoryID:   int(suite.category2.ID),
				VendorID:     suite.vendor.ID,
			},
		},
		{
			title: "Should create service with name make up for vendor2",
			payload: models.ServiceInput{
				Name:         fmt.Sprintf("make up"),
				Price:        utils.Float64ToPointer(40),
				Duration:     40,
				DurationType: "mins",
				Trending:     utils.BoolToPointer(true),
				CategoryID:   int(suite.category3.ID),
				VendorID:     suite.vendor2.ID,
			},
		},
		{
			title: "Should error out if category does not exit",
			payload: models.ServiceInput{
				Name:         fmt.Sprintf("hair cut"),
				Price:        utils.Float64ToPointer(40),
				Duration:     40,
				DurationType: "mins",
				Trending:     utils.BoolToPointer(true),
				CategoryID:   100000,
				VendorID:     suite.vendor.ID,
			},
			error:         true,
			expectedError: "category with id 100000 does not exit",
		},
		{
			title: "Should error out if vendor does not exit",
			payload: models.ServiceInput{
				Name:         fmt.Sprintf("make up"),
				Price:        utils.Float64ToPointer(40),
				Duration:     40,
				DurationType: "mins",
				Trending:     utils.BoolToPointer(true),
				CategoryID:   int(suite.category.ID),
				VendorID:     "cb7jskehsgjwwlsldjkdkd",
			},
			error:         true,
			expectedError: "vendor with id cb7jskehsgjwwlsldjkdkd does not exit",
		},
		{
			title: "Should error out if service already exit for a vendor",
			payload: models.ServiceInput{
				Name:         fmt.Sprintf("make up"),
				Price:        utils.Float64ToPointer(40),
				Duration:     40,
				DurationType: "mins",
				Trending:     utils.BoolToPointer(true),
				CategoryID:   int(suite.category.ID),
				VendorID:     suite.vendor.ID,
			},
			error:         true,
			expectedError: "service with name make up already exit",
		},
	}

	for _, test := range testCases {
		newService, err := categoryService.CreateService(test.payload)

		if test.error {
			suite.NotNil(err)
			suite.Equal(test.expectedError, err.Error())
		} else {
			var service models.Service
			categoryService.DB.First(&service, newService.ID)
			suite.Nil(err)
			suite.Equal(service.ID, newService.ID, test.title)
			suite.Equal(service.Name, newService.Name, test.title)
			suite.Equal(service.CategoryID, newService.CategoryID, test.title)
			suite.Equal(service.VendorID, newService.VendorID, test.title)
		}
	}
}

func (suite *CategoryServiceTestSuite) TestUpdateService() {

	data := models.ServiceInput{
		Name:         fmt.Sprintf("braiding"),
		Price:        utils.Float64ToPointer(40),
		Duration:     40,
		DurationType: "mins",
		Trending:     utils.BoolToPointer(true),
		CategoryID:   int(suite.category.ID),
		VendorID:     suite.vendor.ID,
	}
	createdService := suite.seed.VendorService(1, data)

	testCases := []struct {
		title      string
		payload    models.ServiceInputUpdate
		CategoryID uint64
	}{
		{
			title: "Should update service name from braiding to make up",
			payload: models.ServiceInputUpdate{
				Name:       utils.StringToPointer("make up"),
				Price:      utils.Float64ToPointer(40),
				Duration:   utils.IntToPointer(70),
				Trending:   utils.BoolToPointer(false),
				CategoryID: utils.IntToPointer(int(suite.category2.ID)),
			},
			CategoryID: createdService.ID,
		},
	}

	for _, test := range testCases {
		_, err := categoryService.UpdateService(1, test.payload)
		var service models.Service
		categoryService.DB.Find(&service, "id = ?", test.CategoryID)
		suite.Nil(err)
		suite.Equal(*test.payload.Name, service.Name, test.title)
		suite.Equal(*test.payload.Price, service.Price, test.title)
		suite.Equal(uint(*test.payload.Duration), service.Duration, test.title)
		suite.Equal(*test.payload.Trending, service.Trending, test.title)
		suite.Equal(uint64(*test.payload.CategoryID), service.CategoryID, test.title)
	}
}

func checkSortOrder(services []*models.Service, order models.SortPrice) bool {
	var lastValue float64
	result := true
	for index, service := range services {
		if index == 0 {
			lastValue = service.Price
			continue
		}

		if lastValue > service.Price && order == models.SortPriceLowest {
			return false
		}

		if lastValue < service.Price && order == models.SortPriceLowest {
			lastValue = service.Price
			continue
		}

		if lastValue < service.Price && order == models.SortPriceHighest {
			return false
		}

		if lastValue > service.Price && order == models.SortPriceHighest {
			lastValue = service.Price
			continue
		}
	}

	return result
}

func (suite *CategoryServiceTestSuite) TestSearchService() {
	suite.seed.
		LoadAndSeed("users").
		LoadAndSeed("vendors").
		LoadAndSeed("address").
		LoadAndSeed("categories").
		LoadAndSeed("services").
		LoadAndSeed("vendor_address")

	testCases := []struct {
		title          string
		lng            *float64
		lat            *float64
		name           string
		price          models.SortPrice
		expectedLength int
	}{
		{
			title:          "Should return 9 services with the name matching relax",
			name:           "relax",
			expectedLength: 9,
		},
		{
			title:          "Should return 6 services with the name matching relax when users includes longitude, latitude around agegee",
			lat:            utils.Float64ToPointer(6.6411857604980469),
			lng:            utils.Float64ToPointer(3.3054516315460205),
			name:           "relax",
			expectedLength: 6,
		},
		{
			title:          "Should return 9 services and sort by highest to lowest when sort is set to highest",
			name:           "relax",
			price:          models.SortPriceHighest,
			expectedLength: 9,
		},
		{
			title:          "Should return 9 services and sort by lowest to highest when sort is set to lowest",
			price:          models.SortPriceLowest,
			name:           "relax",
			expectedLength: 9,
		},
		{
			title:          "Should return 1 service when user search for name Gel Manicure",
			name:           "Gel Manicure",
			expectedLength: 1,
		},
		{
			title:          "Search should be case insensitive",
			name:           "gel manicure",
			expectedLength: 1,
		},
	}

	for _, test := range testCases {
		services, err := categoryService.SearchService(
			test.lat, test.lng, test.name, nil, &test.price,
		)

		if test.price != "" {
			suite.True(checkSortOrder(services, test.price))
		}

		suite.Nil(err)
		suite.Equal(len(services), test.expectedLength, test.title)
	}
}

func TestCategoryServiceSuite(t *testing.T) {
	suite.Run(t, new(CategoryServiceTestSuite))
}
