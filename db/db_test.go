package db_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syntax753/fluffy-doodle/db"
	"github.com/syntax753/fluffy-doodle/model"
)

func TestDB(t *testing.T) {
	db, err := db.NewPayDB()
	assert.Nil(t, err)
	db.Close()
}

func TestInitialise(t *testing.T) {
	jsonFile, err := os.Open("initialise/test.json")
	assert.Nil(t, err)
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	assert.Nil(t, err)

	var tx []model.TX
	json.Unmarshal(data, &tx)

	t.Logf("%v", tx)
}
