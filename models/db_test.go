package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/syntax753/fluffy-doodle/models"
)

func TestNewDB(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Process test tx",
			args:    args{"test"},
			wantErr: false,
		},
		{
			name:    "Process prod tx",
			args:    args{"prod"},
			wantErr: false,
		},
		{
			name:    "Unknown env",
			args:    args{"unknown"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.NewDB(tt.args.env)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got.IDMap, "Expected transactions")
			}
		})
	}
}
