package helper

import (
	"cf-proposal/common/logservice"
	"cf-proposal/common/types"
	"cf-proposal/domain/model"
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
	"net/http"
)

func GetShortUrl(longUrl string) string {

	hash := fnv.New32a()
	hash.Write([]byte(longUrl))
	return hex.EncodeToString(hash.Sum(nil))
}

func HandleHttpError(w http.ResponseWriter, r *http.Request, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	response, _ := json.Marshal(model.ErrorDto{
		Error: err.Error(),
	})

	w.Write(response)
	logservice.LogError(http.StatusBadRequest, r.Method, types.Path(r.URL.Path), err)
}
