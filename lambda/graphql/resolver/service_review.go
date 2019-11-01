package resolver

import (
	"context"

	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/shared/models"
)

type serviceReviewResolver struct{*Resolver}

func (mutation mutationResolver) CreateReview(ctx context.Context, input models.ServiceReviewInput) (*models.ServiceReview, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	review, err := service.Datastore.ServiceReviewDB.CreateReview(
		user.ID,
		input.VendorID,
		input.ServiceID,
		input.Text,
		input.Rating,
		input.ParentID,
	)

	return review, err
}

func (mutation mutationResolver) UpdateReview(ctx context.Context, input models.ServiceReviewUpdateInput) (*models.ServiceReview, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	review, err := service.Datastore.ServiceReviewDB.UpdateReview(
		user.ID,
		input.Text,
		input.Rating,
		input.ID,
	)

	return review, err
}

func (query *queryResolver) GetServiceReviews(ctx context.Context, serviceID int) (*models.ServiceReviewWithAverageRating, error) {
	service := config.GetServices(ctx)

	reviews, err := service.Datastore.ServiceReviewDB.GetReviews(
		serviceID,
	)

	var counter int
	var total int
	for _, review := range reviews {
		if review.ParentID == 0 && review.Rating != 0 {
			counter += 1
			total += review.Rating
		}
	}

	var averageRatings float64

	if (counter > 0) {
		averageRatings = float64(total/counter)
	}

	res := &models.ServiceReviewWithAverageRating {
		Reviews: reviews,
		AverageRatings: averageRatings,
	}

	return res, err
}
