package main

import (
	"fmt"
	"log"

	payments "github.com/syntax753/fluffy-doodle/api"
	"github.com/tkanos/gonfig"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type config struct {
	Port int
	Env  string
}

const cfgFile = "config.json"

var (
	conf config
)

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
		r.Mount("/api/payments", payments.Routes(fmt.Sprintf("schema/%v.json", conf.Env)))
	})

	return r
}

func main() {

	conf = config{}
	err := gonfig.GetConf(cfgFile, &conf)
	if err != nil {
		panic(err)
	}

	r := Router()

	walkFunc := func(method string, route string, handler http.Handler, middwares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r))
}
