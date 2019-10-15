package models_test

import (
	"fmt"
	"testing"

	"github.com/lucsky/cuid"
	"github.com/stretchr/testify/suite"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
)

var cartService = models.CartDBService{models.Connect(config.LoadConfig())}

type CartTestSuite struct {
	suite.Suite
	vendor   models.Vendor
	user     models.User
	category models.VendorCategory
	service  models.Service
	product  models.Product
	seed     *seeder.Seeder
}

func (suite *CartTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.seed.Tables = []string{"categories", "services", "products", "vendors", "users"}
	suite.user = suite.seed.SeedUser("", "testcartuser", "testduser@gmail.com", nil)
	suite.vendor = suite.seed.SeedVendor("", suite.user.ID, "testcartvendor")
	suite.category = suite.seed.SeedCategory(suite.vendor.ID, "testcartcategory")
	suite.service = suite.seed.SeedService(suite.vendor.ID, "testcartservice", int(suite.category.ID), 1)
	suite.product = suite.seed.SeedProduct(suite.vendor.ID, "testcartproduct1", cuid.New(), 10)
}

func (suite *CartTestSuite) TearDownTest() {
	suite.seed.Clean()
}

func (suite *CartTestSuite) AfterTest() {
	suite.seed.Truncate("carts")
}

func (suite *CartTestSuite) TestCreateCart() {
	vendor2 := suite.seed.SeedVendor("", suite.user.ID, "testcartvendor77")
	user2 := suite.seed.SeedUser("", "testcartuser21", "tescartuser22@gmail.com", nil)

	testCases := []struct {
		Title       string
		UserID      string
		VendorID    string
		CartType    string
		TypeID      string
		Quantity    int
		Error       string
	}{
		{
			Title:       "Should create service cart",
			VendorID:    suite.vendor.ID,
			CartType:    "service",
			Quantity:    6,
			TypeID:      string(suite.service.ID),
		},
		{
			Title:       "Should create product cart",
			VendorID:    vendor2.ID,
			CartType:    "product",
			Quantity:    3,
			TypeID:      suite.product.ID,
		},
		{
			Error:       "Higher Quantity",
			Title:       "Should not create cart",
			VendorID:    suite.vendor.ID,
			CartType:    "product",
			Quantity:    15,
			TypeID:      suite.product.ID,
		},
	}

	for _, test := range testCases {
		cart, err := cartService.CreateCart(
			user2.ID,
			test.VendorID,
			test.CartType,
			test.TypeID,
			test.Quantity,
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(cart.Quantity, test.Quantity)
			suite.Equal(cart.TypeID, test.TypeID)
			suite.Equal(cart.VendorID, test.VendorID)
		}

		if test.Error == "Higher Quantity" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Quantity is more than the available product. %v products is/are available", suite.product.Available))
		}
	}
}

func (suite *CartTestSuite) TestUpdateCart() {
	user2 := suite.seed.SeedUser("", "testcartuser212", "testcartuser212@gmail.com", nil)
	product := suite.seed.SeedProduct(suite.vendor.ID, "testcartproduct12", cuid.New(), 7)
	cart, _ := cartService.CreateCart(
		user2.ID,
		suite.vendor.ID,
		"service",
		string(suite.service.ID),
		4,
	)
	cart2, _ := cartService.CreateCart(
		user2.ID,
		suite.vendor.ID,
		"product",
		product.ID,
		5,
	)
	testCases := []struct {
		Title       string
		UserID      string
		ID          string
		Quantity    int
		CartType    string
		TypeID      string
		Error       string
	}{
		{
			Title:       "Should update service cart",
			ID:          cart.ID,
			UserID:      user2.ID,
			CartType:    "service",
			Quantity:    6,
			TypeID:      string(cart.TypeID),
		},
		{
			Title:       "Should update product cart",
			ID:          cart2.ID,
			UserID:      user2.ID,
			CartType:    "product",
			Quantity:    6,
			TypeID:      cart2.TypeID,
		},
		{
			Title:       "Should fail to update product cart",
			ID:          cart2.ID,
			UserID:      user2.ID,
			CartType:    "product",
			Quantity:    9,
			TypeID:      cart2.TypeID,
			Error:       "Higher quantity",
		},
	}

	for _, test := range testCases {
		updatedCart, err := cartService.UpdateCart(
			user2.ID,
			test.ID,
			test.Quantity,
			test.CartType,
			test.TypeID,
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(updatedCart.ID, test.ID)
			suite.Equal(updatedCart.Quantity, test.Quantity)
			suite.Equal(updatedCart.TypeID, test.TypeID)
			suite.Equal(updatedCart.Type, test.CartType)
		}

		if test.Error == "Higher quantity" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Quantity is more than the available product. %v products is/are available", product.Available))
		}
	}
}

func (suite *CartTestSuite) TestGetAllCart() {
	user2 := suite.seed.SeedUser("", "testcartuser222", "testcartuser222@gmail.com", nil)
	product := suite.seed.SeedProduct(suite.vendor.ID, "testcartproduct11", cuid.New(), 7)
	product2 := suite.seed.SeedProduct(suite.vendor.ID, "testcartproduct21", cuid.New(), 7)
	cartService.CreateCart(
		suite.user.ID,
		suite.vendor.ID,
		"service",
		string(suite.service.ID),
		4,
	)
	cartService.CreateCart(
		user2.ID,
		suite.vendor.ID,
		"product",
		product.ID,
		5,
	)
	cartService.CreateCart(
		user2.ID,
		suite.vendor.ID,
		"product",
		product2.ID,
		2,
	)
	testCases := []struct {
		Title       string
		UserID      string
		Error       string
		userCase    int
	}{
		{
			Title:       "Should get carts for user",
			UserID:      suite.user.ID,
			userCase:    1,
		},
		{
			Title:       "Should get carts for user 2",
			UserID:      user2.ID,
			userCase:    2,
		},
	}

	for _, test := range testCases {
		carts, err := cartService.GetAllCarts(
			test.UserID,
		)

		if test.userCase == 1 {
			suite.Nil(err)
			suite.Equal(len(carts), 1)
		}
		if test.userCase == 2 {
			suite.Nil(err)
			suite.Equal(len(carts), 2)
		}
	}
}

func (suite *CartTestSuite) TestDeleteCart() {
	user2 := suite.seed.SeedUser("", "testcartuser220", "testcartuser220@gmail.com", nil)
	cart, _ := cartService.CreateCart(
		user2.ID,
		suite.vendor.ID,
		"service",
		string(suite.service.ID),
		4,
	)
	testCases := []struct {
		Title       string
		UserID      string
		ID          string
		Error       string
	}{
		{
			Title:       "Should not delete cart",
			ID:          cart.ID,
			UserID:      suite.user.ID,
			Error:       "Not allowed to delete",
		},
		{
			Title:       "Should delete cart",
			ID:          cart.ID,
			UserID:      user2.ID,
		},
	}

	for _, test := range testCases {
		deletedCart, err := cartService.DeleteCart(
			test.UserID,
			test.ID,
		)

		if test.Error == "Not allowed to delete" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("An error occurred deleting cart"))
			suite.Equal(deletedCart, false)
		}

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(deletedCart, true)
		}
	}
}

func TestCartSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}
