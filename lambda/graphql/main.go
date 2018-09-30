package main

import (
	"encoding/json"
	"log"

	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/config"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/resolver"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/models"
	svc "gitlab.com/synergy-designs/style-blitz/shared/services"
)

// Schema object
var schema *graphql.Schema
var services *svc.Services

func init() {
	f := utils.GetSchema()
	schema = graphql.MustParseSchema(f, &resolver.Resolver{})
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

	response := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
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
