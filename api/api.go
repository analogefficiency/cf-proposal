package api

import (
	"cf-proposal/common/types"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/repo/historyrepository"
	"cf-proposal/domain/repo/statisticsrepository"
	"cf-proposal/domain/repo/urlrepository"
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

const basePath types.Path = "/"
const createpath types.Path = "/create"
const deletepath types.Path = "/delete/{id}"
const statisticspath types.Path = "/{id}"

var urldatastore *datastore.UrlRepo
var urlrepo *urlrepository.Url
var historydatastore *datastore.HistoryRepo
var histrepo *historyrepository.History
var statisticsdatastore *datastore.StatisticsRepo
var statrepo *statisticsrepository.Statistic

// TODO: Figure out how to either generalize middleware handlers; time permitting.
func RedirectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "shortUrl", chi.URLParam(r, "shortUrl"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
