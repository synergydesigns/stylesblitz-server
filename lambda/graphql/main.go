package main

import (
	"net/http/httptest"
	"strings"

	"context"

	"github.com/99designs/gqlgen/handler"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/genql"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/middleware"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/resolver"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GraphqlHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := httptest.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))

	for k, v := range request.Headers {
		r.Header.Add(k, v)
	}

	w := httptest.NewRecorder()

	http := handler.GraphQL(genql.NewExecutableSchema(genql.Config{Resolvers: &resolver.Resolver{}}))

	http.ServeHTTP(w, r.WithContext(ctx))
	return events.APIGatewayProxyResponse{
		Body:       w.Body.String(),
		StatusCode: w.Code,
		Headers:    config.GetHeaders(),
	}, nil
}

func main() {
	lambda.Start(
		middleware.ServiceInitialize(
			middleware.AuthMiddleware(
				GraphqlHandler,
			),
		),
	)
}
