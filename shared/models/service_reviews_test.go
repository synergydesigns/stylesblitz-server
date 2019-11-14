package models_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synergydesigns/stylesblitz-server/shared/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
	"github.com/synergydesigns/stylesblitz-server/shared/seeder"
)

var reviewService = models.ServiceReviewDBService{models.Connect(config.LoadConfig())}

type ReviewTestSuite struct {
	suite.Suite
	vendor   models.Vendor
	user     models.User
	service  models.Service
	category models.VendorCategory
	seed     *seeder.Seeder
}

func (suite *ReviewTestSuite) SetupTest() {
	suite.seed = seeder.New()
	suite.seed.Tables = []string{"services", "vendors", "users"}
	suite.user = suite.seed.SeedUser("", "testreviewuser", "testreviewuser@gmail.com", nil)
	suite.vendor = suite.seed.SeedVendor("", suite.user.ID, "testreviewvendor")
	suite.category = suite.seed.SeedCategory(suite.vendor.ID, "testcartcategory")
	suite.service = suite.seed.SeedService(suite.vendor.ID, "testreviewservice", int(suite.category.ID), 1)
}

func (suite *ReviewTestSuite) TearDownTest() {
	suite.seed.Clean()
}

func (suite *ReviewTestSuite) AfterTest() {
	suite.seed.Truncate("service_reviews")
}

func (suite *ReviewTestSuite) TestCreateReview() {
	testCases := []struct {
		Title       string
		UserID      string
		VendorID    string
		ServiceID   int
		Text        string
		Rating      int
		ParentID    int
		Error       string
	}{
		{
			Title:       "Should create service review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Awesome service",
			Rating:      5,
		},
		{
			Error:       "Invalid rating",
			Title:       "Should not create service review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Awesome service",
			Rating:      6,
		},
		{
			Error:       "Invalid rating 2",
			Title:       "Should not create service review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Awesome service",
			Rating:      0,
		},
	}

	for _, test := range testCases {
		review, err := reviewService.CreateReview(
			test.UserID,
			test.VendorID,
			test.ServiceID,
			test.Text,
			test.Rating,
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(review.Rating, test.Rating)
			suite.Equal(review.Text, test.Text)
			suite.Equal(review.UserID, test.UserID)
		}

		if test.Error == "Invalid rating" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Rating must be between 1 and 5 inclusive"))
		}

		if test.Error == "Invalid rating 2" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Rating must be between 1 and 5 inclusive"))
		}
	}
}

func (suite *ReviewTestSuite) TestUpdateReview() {
	user2 := suite.seed.SeedUser("", "testreviewuser212", "testcartreview212@gmail.com", nil)

	review, _ := reviewService.CreateReview(
		suite.user.ID,
		suite.vendor.ID,
		int(suite.service.ID),
		"A random test",
		4,
	)
	testCases := []struct {
		Title       string
		UserID      string
		VendorID    string
		ServiceID   int
		Text        string
		Rating      int
		ParentID    int
		Error       string
	}{
		{
			Title:       "Should update service review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Update works!",
			Rating:      5,
		},
		{
			Error:       "User cannot update review",
			Title:       "Should not update service review",
			UserID:      user2.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Update works!",
			Rating:      5,
		},
		{
			Error:       "Invalid rating",
			Title:       "Should not update service review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Awesome service",
			Rating:      6,
		},
		{
			Error:       "Invalid rating 2",
			Title:       "Should not update service review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Awesome service",
			Rating:      0,
		},
	}

	for _, test := range testCases {
		updatedReview, err := reviewService.UpdateReview(
			test.UserID,
			test.Text,
			test.Rating,
			int(review.ID),
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(updatedReview.ID, review.ID)
			suite.Equal(updatedReview.Rating, test.Rating)
			suite.Equal(updatedReview.Text, test.Text)
		}

		if test.Error == "User cannot update review" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("An error occurred updating review"))
		}

		if test.Error == "Invalid rating" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Rating must be between 1 and 5 inclusive"))
		}

		if test.Error == "Invalid rating 2" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("Rating must be between 1 and 5 inclusive"))
		}
	}
}

func (suite *ReviewTestSuite) TestCreateReply() {
	user2 := suite.seed.SeedUser("", "testreviewuser212", "testcartreview212@gmail.com", nil)

	review, _ := reviewService.CreateReview(
		suite.user.ID,
		suite.vendor.ID,
		int(suite.service.ID),
		"A random test",
		4,
	)

	reviewReply, _ := reviewService.CreateReply(
		user2.ID,
		suite.vendor.ID,
		int(suite.service.ID),
		"Replying a review",
		int(review.ID),
	)

	testCases := []struct {
		Title       string
		UserID      string
		VendorID    string
		ServiceID   int
		Text        string
		Rating      int
		ParentID    int
		Error       string
	}{
		{
			Title:       "Should reply a review",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Update works!",
			ParentID:    int(review.ID),
		},
		{
			Title:       "Should not create a reply",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   1000000,
			Text:        "Update works!",
			ParentID:    int(review.ID),
			Error:       "Not found service",
		},
		{
			Title:       "Should not allow a reply to reply a reply",
			UserID:      suite.user.ID,
			VendorID:    suite.vendor.ID,
			ServiceID:   int(suite.service.ID),
			Text:        "Update works!",
			ParentID:    int(reviewReply.ID),
			Error:       "Only allow one level deep reply",
		},
	}

	for _, test := range testCases {
		reply, err := reviewService.CreateReply(
			test.UserID,
			test.VendorID,
			test.ServiceID,
			test.Text,
			test.ParentID,
		)

		if test.Error == "" {
			suite.Nil(err)
			suite.Equal(reply.ParentID, int(review.ID))
			suite.Equal(reply.Text, test.Text)
		}

		if test.Error == "Not found service" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("an error occurred creating reply"))
		}

		if test.Error == "Only allow one level deep reply" {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("An error occurred. You cannot reply a reply :)"))
		}
	}
}

func (suite *ReviewTestSuite) TestGetServiceReview() {
	user2 := suite.seed.SeedUser("", "testcartuser222", "testcartuser222@gmail.com", nil)
	vendor := suite.seed.SeedVendor("", user2.ID, "testreviewvendor1")
	category := suite.seed.SeedCategory(vendor.ID, "testcartcategory1")
	service1 := suite.seed.SeedService(suite.vendor.ID, "testreviewservice1", int(category.ID), 1)
	service2 := suite.seed.SeedService(vendor.ID, "testreviewservice", int(category.ID), 1)
	review1, _ := reviewService.CreateReview(
		suite.user.ID,
		suite.vendor.ID,
		int(service1.ID),
		"First review",
		4,
	)
	reviewService.CreateReview(
		suite.user.ID,
		suite.vendor.ID,
		int(service1.ID),
		"Second review",
		5,
	)
	reviewService.CreateReview(
		suite.user.ID,
		suite.vendor.ID,
		int(service1.ID),
		"Third review",
		4,
	)
	review4, _ := reviewService.CreateReview(
		user2.ID,
		vendor.ID,
		int(service2.ID),
		"Fourth review",
		3,
	)

	testCases := []struct {
		Title       string
		ServiceID   int
		Error       string
		userCase    int
	}{
		{
			Title:       "Should get reviews for user",
			ServiceID:   review1.ServiceID,
			userCase:    1,
		},
		{
			Title:       "Should get carts for user 2",
			ServiceID:      review4.ServiceID,
			userCase:    2,
		},
	}

	for _, test := range testCases {
		reviews, err := reviewService.GetReviews(
			test.ServiceID,
		)

		if test.userCase == 1 {
			suite.Nil(err)
			suite.Equal(len(reviews.Reviews), 3)
		}
		if test.userCase == 2 {
			suite.Nil(err)
			suite.Equal(len(reviews.Reviews), 1)
		}
	}
}

func (suite *ReviewTestSuite) TestDeleteReview() {
	user2 := suite.seed.SeedUser("", "testreviewuser212", "testcartreview212@gmail.com", nil)

	review, _ := reviewService.CreateReview(
		user2.ID,
		suite.vendor.ID,
		int(suite.service.ID),
		"A random test",
		4,
	)
	review2, _ := reviewService.CreateReview(
		suite.user.ID,
		suite.vendor.ID,
		int(suite.service.ID),
		"A random test",
		4,
	)
	review2Reply, _ := reviewService.CreateReply(
		user2.ID,
		suite.vendor.ID,
		int(suite.service.ID),
		"Replying a review",
		int(review2.ID),
	)

	testCases := []struct {
		Title       string
		UserID      string
		ID          int
		Error       string
		UserCase    int
	}{
		{
			Title:       "Should not delete review",
			ID:          int(review.ID),
			UserID:      suite.user.ID,
			Error:       "Not allowed to delete",
			UserCase:    1,
		},
		{
			Title:       "Should delete review",
			ID:          int(review.ID),
			UserID:      user2.ID,
			UserCase:    2,
		},
		{
			Title:       "Should delete review with replies",
			ID:          int(review2.ID),
			UserID:      suite.user.ID,
			UserCase:    3,
		},
	}

	for _, test := range testCases {
		deletedReview, err := reviewService.DeleteReview(
			test.UserID,
			test.ID,
		)

		if test.UserCase == 1 {
			suite.NotNil(err)
			suite.Equal(err, fmt.Errorf("An error occurred deleting review"))
			suite.Equal(deletedReview, false)
		}

		if test.UserCase == 2 {
			suite.Nil(err)
			suite.Equal(deletedReview, true)
		}

		if test.UserCase == 3 {
			suite.Nil(err)
			suite.Equal(deletedReview, true)
			var deletedReview models.ServiceReview

			reviewService.DB.Unscoped().Where("id = ?", review2Reply.ParentID).First(&deletedReview)
			suite.Equal(deletedReview.Text, "Review Deleted")
		}
	}
}

func TestReviewSuite(t *testing.T) {
	suite.Run(t, new(ReviewTestSuite))
}
