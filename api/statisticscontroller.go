package api

import (
	"cf-proposal/common/logservice"
	"cf-proposal/common/types"
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
	logservice.LogHttpRequest(http.StatusOK, r.Method, types.Path(r.URL.Path))
}
