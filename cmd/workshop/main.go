package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"workshop/internal/handler"
)

func main() {
	r := chi.NewRouter()
	h := handler.NewHandler()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("shutting down server")
}
