package middleware

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
)

func AuthMiddleware(next handlerFunc) handlerFunc {
	return handlerFunc(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		svc := config.GetServices(ctx)

		authHeader := request.Headers["Authorization"]

		if authHeader == "" {
			return next(ctx, request)
		}

		token := strings.Split(authHeader, " ")

		if len(token) != 2 {
			return next(ctx, request)
		}

		user, err := svc.JWT.DecodeToken(token[1])

		if err != nil {
			return events.APIGatewayProxyResponse{
				Headers:    config.GetHeaders(),
				StatusCode: 401,
				Body:       config.AuthenticationError("Authentication failed: Invalid token"),
			}, nil
		}

		ctx = context.WithValue(ctx, config.CTXKeyuser, user)

		return next(ctx, request)
	})
}
