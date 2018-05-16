package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for graphqli
func Handler(request interface{}) (events.APIGatewayProxyResponse, error) {
	// get working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// read template file
	index, _ := ioutil.ReadFile(path.Join(dir, "graphqli/template/index.html"))

	headers := map[string]string{
		"Content-Type": "text/html",
	}
	return events.APIGatewayProxyResponse{
		Body:       string(index),
		StatusCode: 200,
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
