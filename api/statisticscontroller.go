package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/repo/statisticsrepository"
	"cf-proposal/domain/repo/urlrepository"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type StatisticsController struct{}

func (sc StatisticsController) InitController() {
	statisticsdatastore = datastore.InitStatisticsDatastore(sqlite3helper.DbConn)
	statrepo = statisticsrepository.Init(statisticsdatastore)

	urldatastore = datastore.InitUrlDatastore(sqlite3helper.DbConn)
	urlrepo = urlrepository.Init(urldatastore)
}

func (sc StatisticsController) StatisticsRoutes() chi.Router {
	router := chi.NewRouter()
	router.Route(string(statisticspath), func(r chi.Router) {
		r.Use(IdCtx)
		r.Get("/", sc.HandleGetStatistics)
	})
	return router
}

func (sc StatisticsController) HandleGetStatistics(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	_, err := urlrepo.GetUrl(context.Background(), id)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}
	statistic, err := statrepo.GetStatistic(context.Background(), id)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}

	helper.HandleHttpOk(w, r, statistic, http.StatusOK)
}
