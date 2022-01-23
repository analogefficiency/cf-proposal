package logservice

import (
	"cf-proposal/common/constants/logcategory"
	"cf-proposal/common/types"
	"log"
)

func LogHttpRequest(responseCode string, httpRequest string, path types.Path) {
	log.Printf("[%s] \033[33m%s\033[0m %s %s", logcategory.INFO, responseCode, httpRequest, path)
}

func LogError(responseCode string, httpRequest string, path types.Path, err error) {
	log.Printf("[%s] \033[31m%s\033[0m %s %s %s", logcategory.ERROR, responseCode, httpRequest, path, err.Error())
}
