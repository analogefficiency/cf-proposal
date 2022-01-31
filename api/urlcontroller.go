package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/logservice"
	"cf-proposal/common/types"
	"cf-proposal/domain/model"
	"cf-proposal/domain/service"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type UrlController struct{}

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
		helper.HandleHttpError(w, r, err, 500)
		return
	}

	var url model.UrlDto
	err = json.Unmarshal(body, &url)
	if err != nil {
		helper.HandleHttpError(w, r, err, 422)
		return
	}

	createdUrl, err := service.UrlService{}.CreateUrl(url)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}
	helper.HandleHttpOk(w, r, createdUrl, http.StatusCreated)
}

func (uc UrlController) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	err := service.UrlService{}.DeleteUrl(id)
	if err != nil {
		logservice.LogError(http.StatusBadGateway, r.Method, types.Path(r.URL.Path), err)
	} else {
		logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
	}
}

func (uc UrlController) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.Context().Value("shortUrl").(string)
	longUrl, err := service.UrlService{}.RedirectUrl(shortUrl)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
	} else {
		logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
		http.Redirect(w, r, longUrl, http.StatusFound)
	}
}
