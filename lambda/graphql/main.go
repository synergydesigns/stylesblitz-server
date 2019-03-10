package main

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"strings"

	"context"

	"github.com/99designs/gqlgen/handler"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/genql"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/resolver"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/models"
	svc "github.com/synergydesigns/stylesblitz-server/shared/services"
)

// Schema object
var services *svc.Services

func init() {
	// Schema = graphql.MustParseSchema(schema.String(), &resolver.Resolver{})
	services = svc.New()
}

// GraphqlHandler handles all qraphql queries
func GraphqlHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var params models.GraphqlBody

	// set context
	ctx = context.WithValue(ctx, config.CTXKeyservices, services)

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Printf("Could not decode body errors %v", err)
	}

	http := handler.GraphQL(genql.NewExecutableSchema(genql.Config{Resolvers: &resolver.Resolver{}}))

	r := httptest.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	w := httptest.NewRecorder()

	http.ServeHTTP(w, r.WithContext(ctx))

	return events.APIGatewayProxyResponse{
		Body:       w.Body.String(),
		StatusCode: w.Code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(GraphqlHandler)
}
