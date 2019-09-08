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
	user      models.User
	seed      *seeder.Seeder
}

func (suite *CategoryServiceTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.seed.Tables = []string{"categories", "vendors", "users"}
	suite.user = suite.seed.SeedUser("", "testudser", "testduser@gmail.com", "09099350122")
	suite.vendor = suite.seed.SeedVendor("", suite.user.ID, "testvendor")
	suite.vendor2 = suite.seed.SeedVendor("", suite.user.ID, "testvendor")
	suite.vendor2 = suite.seed.SeedVendor("", suite.user.ID, "testvendor2")
	suite.category = suite.seed.VendorCategory(1, suite.vendor.ID, "native")
	suite.category2 = suite.seed.VendorCategory(2, suite.vendor.ID, "make up")
}

func (suite *CategoryServiceTestSuite) TearDownTest() {
	suite.seed.Clean()
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
		suite.Equal(len(services), test.expectedCount)
	}
}

func TestCategoryServiceSuite(t *testing.T) {
	suite.Run(t, new(CategoryServiceTestSuite))
}
