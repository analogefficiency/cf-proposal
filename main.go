package main

import (
	"cf-proposal/infrastructure/sqlite3helper"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "9000"

	log.Printf("Starting up on http://localhost:%s", port)
	sqlite3helper.InitDb()

	conn, err := sql.Open("sqlite3", "./shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := chi.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
