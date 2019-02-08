package main

import (
	"encoding/json"
	"log"

	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/config"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/resolver"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/models"
	"github.com/synergydesigns/stylesblitz-server/lambda/graphql/schema"
	svc "github.com/synergydesigns/stylesblitz-server/shared/services"
)

// Schema object
var Schema *graphql.Schema
var services *svc.Services

func init() {
	Schema = graphql.MustParseSchema(schema.String(), &resolver.Resolver{})
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

	response := Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	resp, err := json.Marshal(response)
	if err != nil {
		log.Printf("unable to unmarshal response %v", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(GraphqlHandler)
}
