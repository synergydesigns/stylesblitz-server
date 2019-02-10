package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

//ImportJSONDataFromFile function to import json from file to map
func ImportJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
	}
	return true
}

// LoadSchema recursively loads graphql schema
// It takes the part to the schema an loads all
// the schema found in that directory and anyother
// directory within that folder
// @params buf := pointer to a buffer
// @params dir := directory to read from
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
