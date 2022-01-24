package logservice

import (
	"cf-proposal/common/constants/logcategory"
	"cf-proposal/common/types"
	"log"
	"net/http"
)

func LogHttpRequest(responseCode int, httpRequest string, path types.Path) {
	log.Printf("[%s] \033[33m%s\033[0m %s %s", logcategory.INFO, http.StatusText(responseCode), httpRequest, path)
}

func LogError(responseCode int, httpRequest string, path types.Path, err error) {
	log.Printf("[%s] \033[31m%s\033[0m %s %s %s", logcategory.ERROR, http.StatusText(responseCode), httpRequest, path, err.Error())
}
