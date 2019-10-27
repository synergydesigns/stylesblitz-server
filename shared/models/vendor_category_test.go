package models_test

import (
	"fmt"
	"testing"

	"github.com/lucsky/cuid"
	"github.com/stretchr/testify/suite"

	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

var vendorService = models.VendorCategoryDBService{models.Connect(config.LoadConfig())}

type VendorCategoryTestSuite struct {
	suite.Suite
	vendor models.Vendor
	user   models.User
	seed   *seeder.Seeder
}

func (suite *VendorCategoryTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.seed.Tables = []string{"categories", "vendors", "users"}
	suite.seed.Tables = []string{"categories", "vendors", "users", "services"}
	suite.user = suite.seed.SeedUser("", "testuser", "testduser@gmail.com", nil)
	suite.vendor = suite.seed.SeedVendor("", suite.user.ID, "testvendor")
}

func (suite *VendorCategoryTestSuite) TearDownTest() {
	suite.seed.Clean()
}

func (suite *VendorCategoryTestSuite) AfterTest() {
	suite.seed.Truncate("categories")
}

func (suite *VendorCategoryTestSuite) TestGetAllCategoriesByVendorID() {
	for _, value := range utils.MakeRange(1, 21) {
		suite.seed.VendorCategory(
			0,
			suite.vendor.ID,
			fmt.Sprintf("testvendor %d", value),
		)
	}

	testCase := []struct {
		Title    string
		VendorID string
		Expected int
	}{
		{
			Title:    "Should return 20 categories for the vendor",
			VendorID: suite.vendor.ID,
			Expected: 20,
		},
		{
			Title:    "Should an 0 categories if no category is found for the vendor",
			VendorID: cuid.New(),
			Expected: 0,
		},
	}

	for _, test := range testCase {
		categories, err := vendorService.GetAllCategoriesByVendorID(test.VendorID)
		suite.Equal(test.Expected, len(categories), test.Title)
		suite.Nil(err)
	}
}

func (suite *VendorCategoryTestSuite) TestCreateCategory() {
	vendor2 := suite.seed.SeedVendor("", suite.user.ID, "testvendor2")

	testCases := []struct {
		Title       string
		VendorID    string
		Name        string
		Description string
		Error       string
	}{
		{
			Title:       "Should create category",
			VendorID:    suite.vendor.ID,
			Name:        "make up",
			Description: "We provide awesome makeup",
		},
		{
			Title:       "Should create new category category if already exist for another vendor",
			VendorID:    vendor2.ID,
			Name:        "make up",
			Description: "We provide awesome braiding service",
		},
		{
			Title:       "Should return an error if category name already exist for a vendor",
			VendorID:    suite.vendor.ID,
			Name:        "make up",
			Description: "We provide awesome makeup",
			Error:       "DuplicateField",
		},
		{
			Title:       "Should return an error if vendorId does not exit",
			VendorID:    cuid.New(),
			Name:        "braiding",
			Description: "We provide awesome braiding",
			Error:       "ForeignKeyNotExist",
		},
	}

	for _, test := range testCases {
		category, err := vendorService.CreateCategory(
			test.VendorID,
			test.Name,
			test.Description,
		)

		if test.Error == "DuplicateField" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Category with name %s already exit", test.Name))
		}

		if test.Error == "ForeignKeyNotExist" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Vendor with id %s does not exit", test.VendorID))
		}

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(category.Name, test.Name)
			suite.Equal(category.Description, test.Description)
			suite.Equal(category.VendorID, test.VendorID)
		}
	}
}

func (suite *VendorCategoryTestSuite) TestUpdateCategory() {
	category := suite.seed.VendorCategory(1, suite.vendor.ID, "braiding")
	category2 := suite.seed.VendorCategory(2, suite.vendor.ID, "make up")

	testCases := []struct {
		Title               string
		VendorID            string
		ExpectedName        string
		Name                *string
		ExpectedDescription string
		Description         *string
		CategoryID          uint64
		Error               string
	}{
		{
			Title:               "Should update category name from braiding to barbing",
			VendorID:            suite.vendor.ID,
			ExpectedName:        "barbing",
			ExpectedDescription: "",
			Name:                utils.StringToPointer("barbing"),
			CategoryID:          category.ID,
		},
		{
			Title:               "Should update category description \"very good barbing shop\"",
			VendorID:            suite.vendor.ID,
			ExpectedName:        "make up",
			Name:                nil,
			ExpectedDescription: "very good barbing shop",
			Description:         utils.StringToPointer("very good barbing shop"),
			CategoryID:          category2.ID,
		},
	}

	for _, test := range testCases {
		category, err := vendorService.UpdateCategory(
			test.CategoryID,
			test.VendorID,
			test.Name,
			test.Description,
		)

		suite.Nil(err)
		suite.Equal(test.ExpectedName, category.Name, test.Title)
		suite.Equal(test.ExpectedDescription, category.Description, test.Title)
	}

}

func (suite *VendorCategoryTestSuite) TestDeleteCategory() {
	category := suite.seed.VendorCategory(1, suite.vendor.ID, "braiding")

	testCases := []struct {
		Title      string
		CategoryID uint64
	}{
		{
			Title:      "Should delete category",
			CategoryID: category.ID,
		},
	}

	for _, test := range testCases {
		deleted, err := vendorService.DeleteCategory(test.CategoryID)
		var deletedCategory models.VendorCategory
		vendorService.DB.Find(&deletedCategory, "id = ?", test.CategoryID)

		suite.Nil(err)
		suite.True(deleted, test.Title)
		suite.Equal(deletedCategory.ID, uint64(0), test.Title)
	}
}

func TestVendorCategorySuite(t *testing.T) {
	suite.Run(t, new(VendorCategoryTestSuite))
}
