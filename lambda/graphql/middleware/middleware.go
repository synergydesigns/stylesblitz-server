package middleware

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

var services *service.Services

func init() {
	services = service.New()
}

type handlerFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func ServiceInitialize(next handlerFunc) handlerFunc {
	return handlerFunc(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		ctx = context.WithValue(ctx, config.CTXKeyservices, services)

		return next(ctx, request)
	})
}
