// Package payments configures the routes for the payments api
package payments_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/syntax753/fluffy-doodle/model"

	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	payments "github.com/syntax753/fluffy-doodle/api"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func init() {
	r = chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
	)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api/payments", payments.Routes("../schema/test.json"))
	})
}

func TestTX_GetUnknownID(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/payments/123", nil)

	r.ServeHTTP(rec, req)
	expectedStatus := 404

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 404 status")
}

func TestTX_GetSuccess(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)

	r.ServeHTTP(rec, req)
	expectedStatus := 200

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 200 status")
}

func TestTX_GetAll(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/payments", nil)

	r.ServeHTTP(rec, req)
	expectedStatus := 200

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 200 status")
}

func TestTX_Create(t *testing.T) {
	tx := &model.TX{ID: "1234", Type: "Payment"}
	b, err := json.Marshal(tx)

	if err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/payments", bytes.NewBuffer(b))

	r.ServeHTTP(rec, req)
	expectedStatus := 201

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 201 status")
}

func TestTX_CreateNonPayment(t *testing.T) {
	tx := &model.TX{ID: "", Type: "Holiday"}
	b, err := json.Marshal(tx)

	if err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/payments", bytes.NewBuffer(b))

	r.ServeHTTP(rec, req)
	expectedStatus := 400

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 400 status")
}

func TestTX_UpdateMissingType(t *testing.T) {
	tx := &model.TX{ID: "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", OrganisationID: "1245"}
	b, err := json.Marshal(tx)

	if err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/v1/api/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", bytes.NewBuffer(b))

	r.ServeHTTP(rec, req)
	expectedStatus := 400

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 400 status")
}

func TestTX_Update(t *testing.T) {
	tx := &model.TX{Type: "Payment", ID: "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", OrganisationID: "1245"}
	b, err := json.Marshal(tx)

	if err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/v1/api/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", bytes.NewBuffer(b))

	r.ServeHTTP(rec, req)
	expectedStatus := 200

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 200 status")
}

func TestTX_Delete(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/api/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)

	r.ServeHTTP(rec, req)
	expectedStatus := 204

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 204 status")
}

func TestTX_DeleteNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/api/payments/123324", nil)

	r.ServeHTTP(rec, req)
	expectedStatus := 404

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 404 status")
}
