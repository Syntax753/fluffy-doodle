// Package payments configures the routes for the payments api
package payments

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPayment(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}

	tests := []struct {
		Name           string
		Method         string
		URL            string
		ExpectedStatus int
		ExpectedBody   string
		ExpectedError  string
	}{
		{
			"Get single payment",
			"GET",
			"/1",
			200,
			"GET",
			"",
		},
		{
			"Wrong url",
			"GET",
			"/",
			404,
			"",
			"",
		},
	}

	router := Routes()

	for _, tt := range tests {

		r, err := http.NewRequest(tt.Method, tt.URL, nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		assert.Equal(t, tt.ExpectedStatus, w.Code)
		if tt.ExpectedBody != "" {
			assert.Contains(t, w.Body.String(), tt.ExpectedBody)
		}

		if tt.ExpectedError != "" {
			assert.Contains(t, w.Body, tt.ExpectedError)
		}

	}
}
