package api

import (
	"cf-proposal/common/types"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/urlservice"
)

const basePath types.Path = "/"
const createpath types.Path = "/create"

var urlRepo *repository.UrlRepo
var urlService *urlservice.Url
