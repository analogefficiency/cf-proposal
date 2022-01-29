package main

import (
	"cf-proposal/api"
	"cf-proposal/common/logservice"
	"cf-proposal/domain/service"
	"cf-proposal/infrastructure/sqlite3helper"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "9000"

	// Init Database
	logservice.LogInfo(fmt.Sprintf("Starting up on http://localhost:%s", port))
	sqlite3helper.InitDb("shortener")

	// Init Services
	service.UrlService{}.InitService()
	service.StatisticsService{}.InitService()

	// Define Routes
	r := chi.NewRouter()
	r.Route("/{shortUrl}", func(r chi.Router) {
		r.Use(api.RedirectCtx)
		r.Get("/", api.UrlController{}.HandleRedirect)
	})
	r.Mount("/url", api.UrlController{}.UrlRoutes())
	r.Mount("/statistics", api.StatisticsController{}.StatisticsRoutes())

	log.Fatal(http.ListenAndServe(":"+port, r))
}
