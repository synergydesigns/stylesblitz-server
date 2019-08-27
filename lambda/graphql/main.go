package main

import (
	"net/http/httptest"
	"strings"

	"context"

	"github.com/99designs/gqlgen/handler"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/genql"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/resolver"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	svc "github.com/synergydesigns/stylesblitz-server/shared/services"
)

var services *svc.Services

func init() {
	services = svc.New()
}

func GraphqlHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx = context.WithValue(ctx, config.CTXKeyservices, services)

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
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Control-Allow-Credentials":    "true",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET,POST,OPTIONS",
			"Access-Control-Allow-Headers": "Connection, Host, Origin, Referer, Access-Control-Request-Method, Access-Control-Request-Headers, User-Agent, Accept, Content-Type, Authorization, Content-Length, X-Requested-With, Accept-Encoding, Accept-Language",
		},
	}, nil
}

func main() {
	lambda.Start(GraphqlHandler)
}
