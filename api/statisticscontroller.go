package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/service"
	"net/http"

	"github.com/go-chi/chi"
)

type StatisticsController struct{}

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
	statisticsDto, err := service.StatisticsService{}.GetStatistic(id)
	if err != nil {
		helper.HandleHttpError(w, r, err, 400)
		return
	}
	helper.HandleHttpOk(w, r, statisticsDto, http.StatusOK)
}
