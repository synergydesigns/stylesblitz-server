package config

import (
	"encoding/json"
)

type authenticationError struct {
	AuthenticationError string `json:"authenticationError"`
}

type exception struct {
	InvalidArgs []string `json:"invalidArgs,omitempty"`
	Stacktrace  []string `json:"stacktrace,omitempty"`
}

type extensions struct {
	Code      string    `json:"code,omitempty"`
	Exception exception `json:"exception,omitempty"`
}

type errorDetails struct {
	Message    string     `json:"message"`
	Locations  []string   `json:"locations,omitempty"`
	Path       []string   `json:"path,omitempty"`
	Extensions extensions `json:"extentions"`
}

func parseError(err []errorDetails) string {
	resp, _ := json.Marshal(err)

	return string(resp)
}

func AuthenticationError(message string) string {
	errorMessage := []errorDetails{
		{
			Message: message,
			Extensions: extensions{
				Code: "UNAUTHENTICATED",
			},
		},
	}

	return parseError(errorMessage)
}
