package main

import (
	"fmt"
	"log"
	"os"

	payments "github.com/syntax753/fluffy-doodle/api"

	"github.com/BurntSushi/toml"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/boltdb/bolt"
	"github.com/syntax753/fluffy-doodle/db"
)

type config struct {
	Port int
}

const configFile = "config.toml"

var (
	conf  config
	dbPay *bolt.DB
)

func init() {
	log.Println("Starting router")

	_, err := db.NewPayDB()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		log.Fatal("Can't open config file\n", err)
	}

	log.SetOutput(os.Stdout)
}

// Router retuns the mux for all the restful endpoints
func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
	)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api/payments", payments.Routes("prod"))
	})

	return r
}

func main() {
	r := Router()

	walkFunc := func(method string, route string, handler http.Handler, middwares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	r.log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))

}
