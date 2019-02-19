// Package payments configures the routes for the payments api
package payments_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
}

func TestTXGetUnknown(t *testing.T) {
	// Reset
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api/payments", payments.Routes("test"))
	})

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/payments/123", nil)
	r.ServeHTTP(rec, req)

	expectedStatus := 404

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 404 status")
}

func TestTXGetSuccess(t *testing.T) {
	// Reset
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api/payments", payments.Routes("test"))
	})

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)
	r.ServeHTTP(rec, req)

	expectedStatus := 200

	assert.Equal(t, expectedStatus, rec.Code, "Expected a 200 status")
}
