package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var Schema graphql.Schema

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Println(p.Context.Value(currentUser), "===============")
					return p.Context.Value(currentUser), nil
				},
			},
		},
	},
)

func init() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	if err != nil {
		log.Fatal("Failed to create graphql schema, error: %v", err)
	}

	Schema = s
}

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Body struct {
	Query         string                 `json:"query,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
}

// QueryNameNotProvided is thrown when a name is not provided
var (
	ErrorNameNotProvided = errors.New("no query was provided in the HTTP body")
)

const currentUser string = "currentUser"

// Handler endpoint Handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// query := request.QueryStringParameters

	// fmt.Println(query)

	// user := User{1, "Enaho Murphy"}

	var params Body

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body errors %v", err)
	}

	// queryResponse, _ := json.Marshal(request)
	user := User{
		Name: "Enaho Murphy",
		ID:   2,
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: params.Query,
		Context:       context.WithValue(context.Background(), currentUser, user),
	})

	if len(result.Errors) > 0 {
		return events.APIGatewayProxyResponse{}, ErrorNameNotProvided
	}
	type Response struct {
		Data   interface{}                `json:"data"`
		Errors []gqlerrors.FormattedError `json:"errors,omitempty"`
	}

	data := Response{
		Data:   result.Data,
		Errors: result.Errors,
	}
	resp, _ := json.Marshal(data)
	// fmt.Println(result)
	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

//Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}
