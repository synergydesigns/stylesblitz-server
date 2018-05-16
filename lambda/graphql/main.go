package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/models"
	"gitlab.com/synergy-designs/style-blitz/lambda/graphql/schema"
)

// Schema  holds the gra[hql schema object]\
var Schema graphql.Schema

func init() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.QueryType,
	})

	if err != nil {
		log.Fatalf("Failed to create graphql schema, error: %v", err)
	}

	Schema = s
}

// GraphqlHandler handles all qraphql queries
func GraphqlHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var params models.GraphqlBody

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Printf("Could not decode body errors %v", err)
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: params.Query,
	})

	if len(result.Errors) > 0 {
		data := models.GraphqlResponse{
			Data:   result.Data,
			Errors: result.Errors,
		}
		resp, _ := json.Marshal(data)
		return events.APIGatewayProxyResponse{
			Body: string(resp),
		}, nil
	}

	data := models.GraphqlResponse{
		Data:   result.Data,
		Errors: result.Errors,
	}
	resp, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(GraphqlHandler)
}
