package db_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/syntax753/fluffy-doodle/db"
)

func TestTXNotFound_Error(t *testing.T) {
	type fields struct {
		ID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Create error",
			fields: fields{
				"5",
			},
			want: "TX with ID 5 not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &db.TXNotFound{
				ID: tt.fields.ID,
			}
			if got := tx.Error(); got != tt.want {
				t.Errorf("TXNotFound.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTXDB_Find(t *testing.T) {
	data, err := db.ProcessFile("schema/test.json")
	if err != nil {
		log.Fatalf("Can't load transactions from %v: %v\n", "schema/test.json", err)
	}

	txdb := db.TXDB{TXs: *data.AsMap()}

	tests := []struct {
		name     string
		ID       string
		contains string
		wantErr  bool
	}{
		{
			name:     "Transaction not found",
			ID:       "id-not-found",
			contains: "",
			wantErr:  true,
		},
		{
			name:     "Transaction found",
			ID:       "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			contains: "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := txdb.Find(tt.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TXDB.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.ID, got.ID, "Expecting found transaction with same id %v", tt.ID)
			}
		})
	}
}
