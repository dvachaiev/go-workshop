package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/api/jokes"
	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	r := chi.NewRouter()
	h := handler.NewHandler(apiClient)

	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Print("starting server at ", path)
	err = http.ListenAndServe(path, r)
	log.Fatal(err)

	log.Print("shutting down server")
}
