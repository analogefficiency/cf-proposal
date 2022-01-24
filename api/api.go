package api

import (
	"cf-proposal/common/types"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/urlservice"
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

const basePath types.Path = "/"
const createpath types.Path = "/create"

var urlRepo *repository.UrlRepo
var urlService *urlservice.Url

func RedirectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "shortUrl", chi.URLParam(r, "shortUrl"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
