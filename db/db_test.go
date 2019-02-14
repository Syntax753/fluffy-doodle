package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syntax753/fluffy-doodle/db"
)

func TestDB(t *testing.T) {
	db, err := db.NewPayDB()
	assert.Nil(t, err)
	db.Close()
}
