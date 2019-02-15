// Package payments configures the routes for the payments api
package payments

import (
	"log"
	"net/http"

	"github.com/syntax753/fluffy-doodle/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Routes defines the api for the /api group of methods
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetPayment)
	// router.Post("/{id}", handlerFn)
	return router
}

// GetPayment returns a single payment given an id
func GetPayment(w http.ResponseWriter, r *http.Request) {
	IDMap := r.Context().Value("IDMap")
	id := chi.URLParam(r, "id")
	log.Printf("GET payment id %s\n", id)

	payment := &model.TX{}
	render.JSON(w, r, payment)
}
