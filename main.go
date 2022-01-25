package main

import (
	"cf-proposal/api"
	"cf-proposal/infrastructure/sqlite3helper"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "9000"

	log.Printf("Starting up on http://localhost:%s", port)
	sqlite3helper.InitDb("shortener")
	r := chi.NewRouter()
	api.UrlController{}.InitController()
	r.Mount("/url", api.UrlController{}.UrlRoutes())
	r.Route("/{shortUrl}", func(r chi.Router) {
		r.Use(api.RedirectCtx)
		r.Get("/", api.UrlController{}.HandleRedirect)
	})
	r.Mount("/statistics", api.StatisticsController{}.StatisticsRoutes())
	log.Fatal(http.ListenAndServe(":"+port, r))
}
