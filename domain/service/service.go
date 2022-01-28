package service

import (
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/repo/historyrepository"
	"cf-proposal/domain/repo/urlrepository"
)

var urldatastore *datastore.UrlRepo
var urlrepo *urlrepository.Url
var historydatastore *datastore.HistoryRepo
var histrepo *historyrepository.History
