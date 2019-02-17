package models

import (
	"reflect"
	"testing"
)

func TestData_AsMap(t *testing.T) {

	data := IDMap{
		"1": TX{ID: "1"},
		"2": TX{ID: "2"},
	}

	type fields struct {
		TXs []TX
	}
	tests := []struct {
		name   string
		fields fields
		want   IDMap
	}{
		{
			name:   "Correct count",
			fields: fields{TXs: []TX{TX{ID: "1"}, TX{ID: "2"}}},
			want:   data,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &Data{
				TXs: tt.fields.TXs,
			}
			if got := *data.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.AsMap() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
