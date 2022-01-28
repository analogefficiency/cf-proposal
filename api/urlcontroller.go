package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/logservice"
	"cf-proposal/common/messages"
	"cf-proposal/common/types"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/repo/historyrepository"
	"cf-proposal/domain/repo/urlrepository"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type UrlController struct{}

func (uc UrlController) InitController() {
	urldatastore = datastore.InitUrlDatastore(sqlite3helper.DbConn)
	historydatastore = datastore.InitHistoryDatastore(sqlite3helper.DbConn)

	urlrepo = urlrepository.Init(urldatastore)
	histrepo = historyrepository.Init(historydatastore)
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
		helper.HandleHttpError(w, r, err, 500)
		return
	}

	var url datastore.Url
	err = json.Unmarshal(body, &url)
	if err != nil {
		helper.HandleHttpError(w, r, err, 422)
		return
	}

	createdUrl, err := urlrepo.Create(context.Background(), url)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			logservice.LogInfo(fmt.Sprintf(messages.SHORT_URL_EXISTS, url.LongUrl))
			urlDto, err := urlrepo.GetShortUrlByLongUrl(context.Background(), url.LongUrl)
			if err != nil {
				helper.HandleHttpError(w, r, err, 400)
			}
			helper.HandleHttpOk(w, r, urlDto, 200)
		} else {
			helper.HandleHttpError(w, r, err, 400)
		}
		return
	}

	helper.HandleHttpOk(w, r, createdUrl, http.StatusCreated)
}

func (uc UrlController) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	err := urlrepo.DeleteUrl(context.Background(), id)
	if err != nil {
		logservice.LogError(http.StatusBadGateway, r.Method, types.Path(r.URL.Path), err)
	} else {
		logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
	}
}

func (uc UrlController) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.Context().Value("shortUrl").(string)
	data, err := urlrepo.GetLongUrl(context.Background(), shortUrl)
	if err != nil {
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			helper.HandleHttpError(w, r, &helper.CustomError{Message: fmt.Sprintf(messages.SHORT_URL_DOES_NOT_EXIST, shortUrl)}, 400)
			return
		}
		helper.HandleHttpError(w, r, err, 400)
		return
	}
	logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
	histrepo.Insert(context.Background(), data.UrlID)
	http.Redirect(w, r, data.LongUrl, http.StatusFound)
}
