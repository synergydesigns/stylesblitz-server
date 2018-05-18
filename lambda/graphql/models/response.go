package models

import "github.com/graphql-go/graphql/gqlerrors"

// GraphqlResponse struct for graphql
type GraphqlResponse struct {
	Data   interface{}                `json:"data"`
	Errors []gqlerrors.FormattedError `json:"errors,omitempty"`
}

// GraphqlBody Request Body
type GraphqlBody struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}
