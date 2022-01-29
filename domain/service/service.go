package service

import (
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/repo/historyrepository"
	"cf-proposal/domain/repo/statisticsrepository"
	"cf-proposal/domain/repo/urlrepository"
)

var urldatastore *datastore.UrlDatastore
var urlrepo *urlrepository.Url
var historydatastore *datastore.HistoryDatastore
var histrepo *historyrepository.History
var statisticsdatastore *datastore.StatisticsDatastore
var statrepo *statisticsrepository.Statistic
