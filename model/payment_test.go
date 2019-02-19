package model_test

import (
	"testing"

	"github.com/syntax753/fluffy-doodle/model"

	"github.com/stretchr/testify/assert"
)

func TestDB_GetAllTX(t *testing.T) {
	db, err := model.NewDB("../schema/test.json")
	if err != nil {
		t.Fatalf("Could not open database: %v", err)
	}

	txs, err := db.GetAllTX()
	assert.Equal(t, 3, len(txs), "Expected 3 transactions in the test db")
}

func TestDB_GetTX(t *testing.T) {
	db, err := model.NewDB("../schema/test.json")
	if err != nil {
		t.Fatalf("Could not open database: %v", err)
	}

	tx, err := db.GetTX("1111")
	_, ok := err.(*model.TXNotFound)
	assert.True(t, ok, "Expected transaction not found")
	assert.Nil(t, tx, "Expected nil transaction")

	tx, err = db.GetTX("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	assert.NoError(t, err, "Expected transaction found")
	assert.Equal(t, "Payment", tx.Type, "Expected transaction type to be Payment")
}

func TestDB_CreateTX(t *testing.T) {
	db, err := model.NewDB("../schema/test.json")
	if err != nil {
		t.Fatalf("Could not open database: %v", err)
	}

	tx := &model.TX{
		Type: "Car",
	}

	// Type = Payment is mandatory
	tx, err = db.CreateTX(*tx)
	_, ok := err.(*model.TXInvalid)
	assert.True(t, ok, "Expected transaction raised as invalid")
	assert.Nil(t, tx, "Expected nil transaction")

	tx = &model.TX{
		Type: "Payment",
	}

	// Success
	tx.Type = "Payment"
	tx, err = db.CreateTX(*tx)
	assert.NoError(t, err, "Expected transaction to be created")
	assert.NotNil(t, tx, "Expected transaction to be created")

	// Should be able to get now
	tx, err = db.GetTX(tx.ID)
	assert.NoError(t, err, "Expected transaction found")
}

func TestDB_UpdateTX(t *testing.T) {
	db, err := model.NewDB("../schema/test.json")
	if err != nil {
		t.Fatalf("Could not open database: %v", err)
	}

	// First get
	tx, err := db.GetTX("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	assert.NoError(t, err, "Expected transaction found")

	tx.OrganisationID = "12233"
	tx, err = db.UpdateTX(*tx)
	assert.NoError(t, err, "Expected transaction to be updated")
	assert.NotNil(t, tx, "Expected transaction to be updated")
	assert.Equal(t, "12233", tx.OrganisationID, "Expected the organisation to be updated")

	tx.ID = "1234"
	tx, err = db.UpdateTX(*tx)
	_, ok := err.(*model.TXNotFound)
	assert.True(t, ok, "Expected transaction not found")
	assert.Nil(t, tx, "Expected nil transaction")

	// ID is mandatory
	// First get
	tx, err = db.GetTX("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	assert.NoError(t, err, "Expected transaction found")

	tx.ID = ""
	tx, err = db.UpdateTX(*tx)
	_, ok = err.(*model.TXInvalid)
	assert.True(t, ok, "Expected transaction raised as invalid")
	assert.Nil(t, tx, "Expected nil transaction")
}

func TestDB_DeleteTX(t *testing.T) {
	db, err := model.NewDB("../schema/test.json")
	if err != nil {
		t.Fatalf("Could not open database: %v", err)
	}

	err = db.DeleteTX("1234")
	_, ok := err.(*model.TXNotFound)
	assert.True(t, ok, "Expected transaction not found")

	err = db.DeleteTX("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	assert.NoError(t, err, "Expected deletion to be successful")

	// Fetch all and count
	txs, err := db.GetAllTX()
	assert.Equal(t, 2, len(txs), "Expected 2 transactions left in the test db")
}
