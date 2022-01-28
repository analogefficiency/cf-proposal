package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/statisticsservice"
	"cf-proposal/domain/services/urlservice"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type StatisticsController struct{}

func (sc StatisticsController) InitController() {
	statisticsRepo = repository.InitStatisticsRepo(sqlite3helper.DbConn)
	statisticsService = statisticsservice.Init(statisticsRepo)

	urlRepo = repository.InitUrlRepo(sqlite3helper.DbConn)
	urlService = urlservice.Init(urlRepo)
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
	_, err := urlService.GetUrl(context.Background(), id)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}
	statistic, err := statisticsService.GetStatistic(context.Background(), id)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}

	helper.HandleHttpOk(w, r, statistic, http.StatusOK)
}
