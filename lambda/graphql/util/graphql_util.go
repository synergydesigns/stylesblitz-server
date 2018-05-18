package util

import (
	"bytes"
	"log"
	"os"
	"path"

	"gitlab.com/synergy-designs/style-blitz/shared"
)

// GetSchema gets the schema
func GetSchema() string {
	dir, _ := os.Getwd()
	schemaBytes := &bytes.Buffer{}

	loadError := shared.LoadSchema(schemaBytes, path.Join(dir, "schema"))

	if loadError != nil {
		log.Println(loadError)
		return schemaBytes.String()
	}
	return schemaBytes.String()
}
