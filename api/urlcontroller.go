package api

import (
	"cf-proposal/common/logservice"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/urlservice"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type UrlController struct{}

func (uc UrlController) Routes() chi.Router {
	router := chi.NewRouter()
	router.Post(string(createpath), uc.HandleCreate)

	// Init repository and service for URL entity
	urlRepo = repository.InitUrlRepo(sqlite3helper.DbConn)
	urlService = urlservice.Init(urlRepo)
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
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	createdUrl, err := urlService.Create(context.Background(), url)

	if err != nil {
		w.WriteHeader(400)
		logservice.LogError("400", "GET", createpath, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(createdUrl)
	w.Write(jsonResp)
	logservice.LogHttpRequest("200", "POST", createpath)
}
