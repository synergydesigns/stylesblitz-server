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
	)

	return review, err
}

func (mutation mutationResolver) CreateReply(ctx context.Context, input models.ServiceReviewInput) (*models.ServiceReview, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	reply, err := service.Datastore.ServiceReviewDB.CreateReply(
		user.ID,
		input.VendorID,
		input.ServiceID,
		input.Text,
		input.ParentID,
	)

	return reply, err
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

	return reviews, err
}

func (mutation mutationResolver) DeleteReview(ctx context.Context, id int) (*bool, error) {
	service := config.GetServices(ctx)
	user := config.GetUser(ctx)

	deleted, err := service.Datastore.ServiceReviewDB.DeleteReview(
		user.ID,
		id,
	)

	return &deleted, err
}
