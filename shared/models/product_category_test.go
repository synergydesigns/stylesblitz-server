package models_test

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/suite"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
)

var productCategoryService = models.ProductCategoryDBService{models.Connect(config.LoadConfig())}

type ProductCategoryTestSuite struct {
	suite.Suite
	vendor models.Vendor
	shop   models.Shop
	asset  models.Asset
	seed   *seeder.Seeder
}

func (suite *ProductCategoryTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.seed.Tables = []string{"assets", "shops", "vendors", "users"}
	user := suite.seed.SeedUser("", "testproductcategoryuser", "testproductcategoryuser@gmail.com", nil)
	suite.vendor = suite.seed.SeedVendor("", user.ID, "testproductcategory")
	suite.shop = suite.seed.SeedShop("", suite.vendor.ID, "shop1")
}

func (suite *ProductCategoryTestSuite) TearDownTest() {
	suite.seed.Clean()
}

func (suite *ProductCategoryTestSuite) AfterTest() {
	suite.seed.Truncate("product_categories")
}

func (suite *ProductCategoryTestSuite) TestCreateCategory() {
	testCases := []struct {
		Title       string
		AssetID     string
		VendorID    string
		ShopID      string
		Name        string
		Description string
		ParentID    string
		Error       string
	}{
		{
			Title:       "Should create product category",
			AssetID:     "",
			VendorID:    suite.vendor.ID,
			ShopID:      suite.shop.ID,
			Name:        "my product",
			Description: "A new product category",
		},
	}

	for _, test := range testCases {
		category, err := productCategoryService.CreateCategory(
			test.VendorID,
			test.ShopID,
			test.AssetID,
			test.Name,
			test.Description,
			"",
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(category.Name, test.Name)
			suite.Equal(category.Description, test.Description)
			suite.Equal(category.VendorID, test.VendorID)
		}
	}
}

func (suite *ProductCategoryTestSuite) TestUpdateCategory() {
	createdCategory, _ := productCategoryService.CreateCategory(
		suite.vendor.ID,
		suite.shop.ID,
		"",
		"category test",
		"No description",
		"",
	)

	testCases := []struct {
		Title       string
		ID          string
		AssetID     string
		VendorID    string
		ShopID      string
		Name        string
		Description string
		ParentID    string
		Error       string
	}{
		{
			Title:       "Should update product category",
			ID:          createdCategory.ID,
			AssetID:     "",
			VendorID:    suite.vendor.ID,
			ShopID:      suite.shop.ID,
			Name:        "My updated product",
			Description: "update product category",
		},
	}

	for _, test := range testCases {
		category, err := productCategoryService.UpdateCategory(
			test.ID,
			test.VendorID,
			test.ShopID,
			test.AssetID,
			test.Name,
			test.Description,
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(category.Name, test.Name)
			suite.NotEqual(category.Name, createdCategory.Name)
			suite.Equal(category.Description, test.Description)
			suite.Equal(category.VendorID, test.VendorID)
		}
	}
}

func (suite *ProductCategoryTestSuite) TestDeleteProductCategory() {
	user2 := suite.seed.SeedUser("", "testcategoryuser212", "testcategory212@gmail.com", nil)
	vendor2 := suite.seed.SeedVendor("", user2.ID, "testproductcategory2")

	createdCategory, _ := productCategoryService.CreateCategory(
		suite.vendor.ID,
		suite.shop.ID,
		"",
		"category test",
		"No description",
		"",
	)

	testCases := []struct {
		Title       string
		ID          string
		VendorID    string
		Error       string
	}{
		{
			Title:       "Should not delete product category",
			ID:          createdCategory.ID,
			VendorID:    vendor2.ID,
			Error:       "invalid vendor",
		},
		{
			Title:       "Should delete product category",
			ID:          createdCategory.ID,
			VendorID:    suite.vendor.ID,
		},
	}

	for _, test := range testCases {
		deletedCategory, err := productCategoryService.DeleteCategory(
			test.ID,
			test.VendorID,
		)
		fmt.Println(err, deletedCategory)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(deletedCategory, true)
		}

		if test.Error == "invalid vendor" {
			suite.NotNil(err)
			suite.Equal(false, deletedCategory)
		}
	}
}

func (suite *ProductCategoryTestSuite) TestGetAllProductCategory() {
	shop2 := suite.seed.SeedShop("", suite.vendor.ID, "shop2")

	createdCategory1, _ := productCategoryService.CreateCategory(
		suite.vendor.ID,
		suite.shop.ID,
		"",
		"category test",
		"No description",
		"",
	)

	productCategoryService.CreateCategory(
		suite.vendor.ID,
		suite.shop.ID,
		"",
		"category test2",
		"Whatever",
		"",
	)

	productCategoryService.CreateCategory(
		suite.vendor.ID,
		shop2.ID,
		"",
		"category test2",
		"Whatever",
		"",
	)

	testCases := []struct {
		Title       string
		ShopID      string
		VendorID    string
		Case        int
		Error       string
	}{
		{
			Title:       "Should get all product category for shop 1",
			ShopID:      createdCategory1.ShopID,
			VendorID:    suite.vendor.ID,
			Case:        1,
		},
		{
			Title:       "Should get all product category for shop 2",
			ShopID:      shop2.ID,
			VendorID:    suite.vendor.ID,
			Case:        2,
		},
	}

	for _, test := range testCases {
		categories, err := productCategoryService.GetAllCategories(
			test.VendorID,
			test.ShopID,
		)
		fmt.Println(err, categories)

		if test.Case == 1 {
			suite.Nil(err)
			suite.Equal(2, len(categories))
		}

		if test.Case == 2 {
			suite.Nil(err)
			suite.Equal(1, len(categories))
		}
	}
}

func TestProductCategorySuite(t *testing.T) {
	suite.Run(t, new(ProductCategoryTestSuite))
}
