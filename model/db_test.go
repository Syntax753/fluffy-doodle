package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syntax753/fluffy-doodle/model"
)

func TestNewTestDB(t *testing.T) {
	_, err := model.NewDB("../schema/test.json")
	assert.NoError(t, err, "Corrupt test db")
}

func TestNewProdDB(t *testing.T) {
	_, err := model.NewDB("../schema/prod.json")
	assert.NoError(t, err, "Corrupt prod db")
}
