package api

import (
	"cf-proposal/domain/model"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/urlservice"
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type UrlController struct{}

func (uc UrlController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/create", uc.HandleCreate)

	return router
}

func (uc UrlController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var url model.UrlDto
	err = json.Unmarshal(body, &url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	conn, err := sql.Open("sqlite3", "./shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	urlrepo := repository.InitUrlRepo(conn)
	urlservice := urlservice.Init(urlrepo)
	urlservice.Create(context.Background(), url)
	conn.Close()
}
