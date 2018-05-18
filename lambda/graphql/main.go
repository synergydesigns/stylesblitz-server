package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/resolver"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/models"
)

// Schema object
var Schema *graphql.Schema

func init() {
	Schema = graphql.MustParseSchema(util.GetSchema(), &resolver.Resolver{})
}

// GraphqlHandler handles all qraphql queries
func GraphqlHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var params models.GraphqlBody

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Printf("Could not decode body errors %v", err)
	}
	fmt.Println(params.Query, params.OperationName, params.Variables)

	response := Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err, "===========")
	}

	fmt.Println(string(resp))

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(GraphqlHandler)
}
