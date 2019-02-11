package api

import (
	"net/http"

	"github.com/syntax753/fluffy-doodle/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Routes defines the api for the /api group of methods
func Routes() *chi.Mux {
	// GetPayment is part of CRUD operations to fetch a payment

}

// GetPayment returns a single payment given an id
func GetPayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	Payment := &model.Payment{}

	render.JSON(w, r, "Fetch ID"+id)
}
