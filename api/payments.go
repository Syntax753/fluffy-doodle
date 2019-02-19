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
		log.Fatalf("Cannot open database: %v", err)
	}

	env := &Env{db}

	router := chi.NewRouter()
	router.Get("/{id}", env.GetPayment)
	// router.Post("/{id}", handlerFn)
	return router
}

// GetPayment returns a single payment given an id
func (env *Env) GetPayment(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	log.Printf("GET payment id %s\n", ID)

	tx, err := env.db.GetTX(ID)

	if err != nil {
		render.Status(r, http.StatusNotFound)
		return
	}
	render.JSON(w, r, tx)
}

// CreatePayment returns a single payment given an id
func (env *Env) CreatePayment(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	log.Printf("GET payment id %s\n", ID)

	tx, err := env.db.GetTX(ID)

	if err != nil {
		render.Status(r, http.StatusNotFound)
		return
	}
	render.JSON(w, r, tx)
}
