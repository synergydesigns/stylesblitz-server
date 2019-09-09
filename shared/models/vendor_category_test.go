package models_test

import (
	"fmt"
	"testing"

	"github.com/lucsky/cuid"
	"github.com/stretchr/testify/assert"

	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

var vendorService = models.VendorCategoryDBService{models.Connect(config.LoadConfig())}
var seed = seeder.New()
var user models.User
var vendor models.Vendor

func init() {
	seed.Tables = []string{"categories", "vendors", "users"}
	user = seed.SeedUser("", "testudser", "testduser@gmail.com", "09099350122")
	vendor = seed.SeedVendor("", user.ID, "testvendor")
}

func TestGetAllCategoriesByVendorID(t *testing.T) {
	for _, value := range utils.MakeRange(1, 21) {
		seed.VendorCategory(
			0,
			vendor.ID,
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
			VendorID: vendor.ID,
			Expected: 20,
		},
		{
			Title:    "Should an 0 categories if no category is found for the vendor",
			VendorID: cuid.New(),
			Expected: 0,
		},
	}

	for _, test := range testCase {
		categories, error := vendorService.GetAllCategoriesByVendorID(test.VendorID)
		assert.Equal(t, test.Expected, len(categories), test.Title)
		assert.Nil(t, error)
	}

	seed.Truncate("categories")
}

func TestCreateCategory(t *testing.T) {
	vendor2 := seed.SeedVendor("", user.ID, "testvendor2")

	testCases := []struct {
		Title       string
		VendorID    string
		Name        string
		Description string
		Error       string
	}{
		{
			Title:       "Should create category",
			VendorID:    vendor.ID,
			Name:        "Make up",
			Description: "We provide awesome makeup",
		},
		{
			Title:       "Should create new category category if already exist for another vendor",
			VendorID:    vendor2.ID,
			Name:        "Make up",
			Description: "We provide awesome braiding service",
		},
		{
			Title:       "Should return an error if category name already exist for a vendor",
			VendorID:    vendor.ID,
			Name:        "Make up",
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
			assert.NotNil(t, err)
			assert.Equal(t, err, fmt.Errorf("Category with name %s already exit", test.Name))
		}

		if test.Error == "ForeignKeyNotExist" {
			assert.NotNil(t, err)
			assert.Equal(t, err, fmt.Errorf("Vendor with id %s does not exit", test.VendorID))
		}

		if test.Error == "" {
			assert.Nil(t, err)
			assert.Equal(t, category.Name, test.Name)
			assert.Equal(t, category.Description, test.Description)
			assert.Equal(t, category.VendorID, test.VendorID)
		}
	}

	seed.Truncate("categories")
}

func TestUpdateCategory(t *testing.T) {
	category := seed.VendorCategory(1, vendor.ID, "braiding")
	category2 := seed.VendorCategory(2, vendor.ID, "make up")

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
			VendorID:            vendor.ID,
			ExpectedName:        "barbing",
			ExpectedDescription: "",
			Name:                utils.StringToPointer("barbing"),
			CategoryID:          category.ID,
		},
		{
			Title:               "Should update category description \"very good barbing shop\"",
			VendorID:            vendor.ID,
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

		assert.Nil(t, err)
		assert.Equal(t, test.ExpectedName, category.Name, test.Title)
		assert.Equal(t, test.ExpectedDescription, category.Description, test.Title)
	}

	seed.Truncate("categories")
}

func TestDeleteCategory(t *testing.T) {
	category := seed.VendorCategory(1, vendor.ID, "braiding")

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

		assert.Nil(t, err)
		assert.True(t, deleted, test.Title)
		assert.Equal(t, deletedCategory.ID, uint64(0), test.Title)
	}
	seed.Clean()
}
