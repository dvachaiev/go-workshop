package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	h := handler.NewHandler()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err = http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("shutting down server")
}
