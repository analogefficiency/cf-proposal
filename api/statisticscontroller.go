package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/statisticsservice"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type StatisticsController struct{}

func (sc StatisticsController) InitController() {
	statisticsRepo = repository.InitStatisticsRepo(sqlite3helper.DbConn)

	statisticsService = statisticsservice.Init(statisticsRepo)
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
	data, serviceErr := statisticsService.GetStatistic(context.Background(), id)
	if serviceErr != nil {
		helper.HandleHttpError(w, r, serviceErr, 400)
		return
	}

	helper.HandleHttpOk(w, r, data)
}
