// Package model contains the data level logic and entities
package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Datastore interface abstracts the db layer
type Datastore interface {
	GetAllTX() ([]*TX, error)
	GetTX(string) (*TX, error)
	CreateTX(TX) (*TX, error)
	UpdateTX(TX) (*TX, error)
	DeleteTX(string) error
}

// IDMap represents transactions as a map with id keys
type IDMap map[string]*TX

// DB represents a custom database in this case
// which holds the map of transactions with id keys
// TODO replace with proper sql.DB database
// TODO could use sync.Map depending on benchmark
type DB struct {
	sync.Mutex
	TXs IDMap
}

// NewDB is the constructor for the db layer
func NewDB(jsonFile string) (*DB, error) {
	data, err := processFile(jsonFile)

	return &DB{TXs: data.asMap()}, err
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

// AsMap provides the transaction data as a map with id keys
func (data *Data) asMap() IDMap {
	m := make(IDMap, len(data.TXs))

	for _, v := range data.TXs {
		m[v.ID] = &v
	}

	return m
}
