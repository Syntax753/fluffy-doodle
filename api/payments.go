// Package payments configures the routes for the payments api
package payments

import (
	"log"
	"net/http"

	"github.com/syntax753/fluffy-doodle/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Env holds the db interfaceallowing it to be mocked in tests
type Env struct {
	db model.Datastore
}

// Routes defines the api for the /api group of methods
func Routes(environment string) *chi.Mux {

	db, err := model.NewDB(environment)

	if err != nil {
		log.Fatalf("Cannot open database: %v\n", err)
	}

	env := &Env{db}

	router := chi.NewRouter()
	router.Get("/", env.GetPayments)
	router.Get("/{id}", env.GetPayment)
	router.Post("/", env.CreatePayment)
	router.Put("/{id}", env.UpdatePayment)
	router.Delete("/{id}", env.DeletePayment)
	return router
}

// GetPayments returns all payments from the database
func (env *Env) GetPayments(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET payments")

	txs, err := env.db.GetAllTX()

	if err != nil {
		log.Printf("Error retrieving payments: %v\n", err)
		render.Status(r, http.StatusNotFound)
		return
	}
	render.JSON(w, r, txs)
}

// GetPayment returns a single payment given an id
func (env *Env) GetPayment(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	log.Printf("GET payment id %s\n", ID)

	tx, err := env.db.GetTX(ID)

	if err != nil {
		log.Printf("Error retrieving payment %s: %v\n", ID, err)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, tx)
}

// CreatePayment creates the payment
func (env *Env) CreatePayment(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST payment")

	tx := &model.TX{}
	err := render.DecodeJSON(r.Body, &tx)

	if err != nil {
		log.Printf("Error deserialising payment: %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	tx, err = env.db.CreateTX(*tx)
	if err != nil {
		log.Printf("Error creating payment: %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdatePayment updates the payment
func (env *Env) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	log.Printf("PUT payment id %s\n", ID)

	tx := &model.TX{}
	err := render.DecodeJSON(r.Body, &tx)

	if err != nil {
		log.Printf("Error deserialising payment: %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	// First try and update
	tx, err = env.db.UpdateTX(*tx)

	// Update success
	if err == nil {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, tx)
		return
	}

	// It's ok not to find the transaction to update
	if _, ok := err.(*model.TXNotFound); !ok {
		log.Printf("Error updating payment: %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	// In which case create
	tx, err = env.db.CreateTX(*tx)
	if err != nil {
		log.Printf("Error creating payment: %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeletePayment deletes a single payment given an id
func (env *Env) DeletePayment(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	log.Printf("DELETE payment id %s\n", ID)

	err := env.db.DeleteTX(ID)

	if err != nil {
		log.Printf("Error retrieving payment %s: %v\n", ID, err)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
