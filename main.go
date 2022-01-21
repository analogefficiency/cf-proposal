package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := "9000"

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
