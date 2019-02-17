package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Datastore abstracts the db layer
type Datastore interface {
	GetTX(string) (TX, error)
	CreateTX(TX) (TX, error)
	UpdateTX(TX) (TX, error)
	DeleteTX(string) error
}

// DB represents the actual data access
type DB struct {
	*IDMap
}

// NewDB is the constructor for the db layer
func NewDB(env string) (*DB, error) {
	data, err := processFile(fmt.Sprintf("../schema/%v.json", env))

	return &DB{data.AsMap()}, err
}

// ProcessFile is a helper func to unmarshal the json transactions
func processFile(file string) (data Data, err error) {
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

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return
	}

	log.Printf("Found %v transactions", len(data.TXs))
	return data, nil
}
