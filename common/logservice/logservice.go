package logservice

import (
	"cf-proposal/common/constants/logcategory"
	"cf-proposal/common/types"
	"log"
)

func LogHttpRequest(category logcategory.Category, responseCode string, path types.Path) {
	log.Printf("[%s] \033[33m%s\033[0m %s", category, responseCode, path)
}
