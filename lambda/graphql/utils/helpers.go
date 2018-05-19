package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// GetSchema gets all the graphql schemas
func GetSchema() string {
	dir, _ := os.Getwd()
	schemaBytes := &bytes.Buffer{}

	loadError := LoadSchema(schemaBytes, path.Join(dir, "graphql/schema"))

	if loadError != nil {
		log.Println(loadError)
		return schemaBytes.String()
	}
	return schemaBytes.String()
}

// LoadSchema we need to be able to parse the graphql schema
// ans load it before execution. This little util dynamically
// loads all graphql file found in a specified folder and
// convert them all to string.
// NOTE it recursively check if there are more folders and
// does the same function.
// @param: buf bytes buffers to write to
// @param: dir directory to read from
// @TODO: make sure function does not load none gql files
func LoadSchema(buf *bytes.Buffer, dir string) error {
	files, error := ioutil.ReadDir(dir)
	if error != nil {
		log.Printf("An error occurred reading file: %v", error)
		return error
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		if !file.IsDir() {
			fileBytes, _ := ioutil.ReadFile(filePath)

			buf.WriteByte('\n')
			buf.Write(fileBytes)
		} else {
			LoadSchema(buf, filePath)
		}
	}

	return nil
}
