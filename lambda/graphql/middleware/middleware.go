package middleware

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	globals "github.com/synergydesigns/stylesblitz-server/shared/config"
	service "github.com/synergydesigns/stylesblitz-server/shared/services"
)

var services *service.Services
var appConfig *globals.Config

func init() {
	appConfig = globals.LoadConfig()
	services = service.New(appConfig)
}

type handlerFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func ServiceInitialize(next handlerFunc) handlerFunc {
	return handlerFunc(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		if request.HTTPMethod == "OPTIONS" {
			return events.APIGatewayProxyResponse{
				StatusCode: 204,
				Headers:    config.GetHeaders(),
			}, nil
		}

		ctx = context.WithValue(ctx, config.CTXKeyservices, services)
		ctx = context.WithValue(ctx, config.CTConfig, appConfig)

		return next(ctx, request)
	})
}
