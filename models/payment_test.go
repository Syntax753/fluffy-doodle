package models_test

import (
	"github.com/syntax753/fluffy-doodle/models"
)

type mockDB struct{}

func (mdb *mockDB) GetTX(ID string) (*models.TX, error) {
	tx := &models.TX{}
	return tx, nil
}

// func TestData_AsMap(t *testing.T) {
// 	txs := (*mockDB).Init()

// }
