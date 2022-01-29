package service

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repo/statisticsrepository"
	"cf-proposal/domain/repo/urlrepository"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"fmt"
	"strconv"
)

type StatisticsService struct{}

func (ss StatisticsService) InitService() {
	statisticsdatastore = datastore.InitStatisticsDatastore(sqlite3helper.DbConn)
	statrepo = statisticsrepository.Init(statisticsdatastore)

	urldatastore = datastore.InitUrlDatastore(sqlite3helper.DbConn)
	urlrepo = urlrepository.Init(urldatastore)
}

func (ss StatisticsService) GetStatistic(id string) (model.StatisticsDto, error) {
	_, err := urlrepo.GetUrl(context.Background(), id)
	if err != nil {
		return model.StatisticsDto{}, err
	}

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return model.StatisticsDto{}, &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	statistic, err := statrepo.GetStatistic(context.Background(), convertedId)
	if err != nil {
		return model.StatisticsDto{}, err
	}
	return model.StatisticsDto{
		UrlID:           statistic.UrlID,
		TwentyFourHours: statistic.TwentyFourHours,
		LastSevenDays:   statistic.LastSevenDays,
		AllTime:         statistic.AllTime,
	}, nil
}
