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

func TestGetAllCategoriesByVendorID(t *testing.T) {
	seed.Tables = []string{"categories", "vendors", "users"}
	user := seed.SeedUser("", "testudser", "testduser@gmail.com", "09099350122")
	vendor := seed.SeedVendor("", user.ID, "testvendor")
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
		t.Log(seed.Tables)
		categories, error := vendorService.GetAllCategoriesByVendorID(test.VendorID)
		assert.Equal(t, test.Expected, len(categories), test.Title)
		assert.Nil(t, error)
	}

	seed.Clean()
}
