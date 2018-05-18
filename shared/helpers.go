package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

//ImportJSONDataFromFile function to import json from file to map
func ImportJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
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

// LoadSchema recursively loads graphql schema
// It takes the part to the schema an loads all
// the schema found in that directory and anyother
// directory within that folder
// @params buf := pointer to a buffer
// @params dir := directory to read from
func LoadSchema(buf *bytes.Buffer, dir string) error {

	files, error := ioutil.ReadDir(dir)
	if error != nil {
		return error
		log.Printf("An error occurred reading file: %v", error)
	}

	for _, file := range files {
		fmt.Println(file.Name())
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
