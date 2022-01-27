package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/logservice"
	"cf-proposal/common/types"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/historyservice"
	"cf-proposal/domain/services/urlservice"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type UrlController struct{}

func (uc UrlController) InitController() {
	urlRepo = repository.InitUrlRepo(sqlite3helper.DbConn)
	historyRepo = repository.InitHistoryRepo(sqlite3helper.DbConn)

	urlService = urlservice.Init(urlRepo)
	historyService = historyservice.Init(historyRepo)
}

func (uc UrlController) UrlRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post(string(createpath), uc.HandleCreate)
	r.Route(string(deletepath), func(r chi.Router) {
		r.Use(IdCtx)
		r.Delete("/", uc.HandleDelete)
	})
	return r
}

// Add Handlers below here, in alpha order

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
		logservice.LogError(http.StatusBadRequest, r.Method, createpath, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(createdUrl)
	w.Write(jsonResp)
	logservice.LogHttpRequest(http.StatusOK, r.Method, createpath)
}

func (uc UrlController) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	err := urlService.DeleteUrl(context.Background(), id)
	if err != nil {
		logservice.LogError(http.StatusBadGateway, r.Method, types.Path(r.URL.Path), err)
	} else {
		logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
	}
}

func (uc UrlController) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.Context().Value("shortUrl").(string)
	data, err := urlService.GetLongUrl(context.Background(), shortUrl)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}
	logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
	historyService.Insert(context.Background(), data.UrlID)
	http.Redirect(w, r, data.LongUrl, http.StatusFound)
}
