package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareFileForLoad(t *testing.T) {
	txs, err := PrepareFileForLoad("schema/test.json")

	assert.Nil(t, err)
	assert.NotEmpty(t, txs, txs)
}
