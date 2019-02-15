package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/syntax753/fluffy-doodle/model"
)

// PrepareForLoad is a helper func to unmarshal json into an array of transactions
func prepareForLoad(jsonData []byte) (txs []model.TX, err error) {
	var data model.Data
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return
	}

	log.Printf("Converted %v transactions", len(data.TXs))

	return data.TXs, nil
}

// PrepareFileForLoad is a helper func to unmarshal a json file into an array of transactions
func PrepareFileForLoad(file string) (txs []model.TX, err error) {
	log.Printf("Converting file %v", file)
	jsonFile, err := os.Open(file)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	return prepareForLoad(jsonData)
}
