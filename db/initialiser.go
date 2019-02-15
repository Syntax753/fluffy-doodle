package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/syntax753/fluffy-doodle/model"
)

// UnmarshalData is a helper func to unmarshal json into memory
func unmarshalData(jsonData []byte) (data model.Data, err error) {
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return
	}

	log.Printf("Found %v transactions", len(data.TXs))
	return data, nil
}

// ProcessFile is a helper func to unmarshal the json transactions
func ProcessFile(file string) (data model.Data, err error) {
	log.Printf("Processing file %v", file)
	jsonFile, err := os.Open(file)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	return unmarshalData(jsonData)
}
