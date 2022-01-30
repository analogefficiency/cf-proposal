package service

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/logservice"
	"cf-proposal/common/messages"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repo/historyrepository"
	"cf-proposal/domain/repo/urlrepository"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

var UrlService UrlServiceFunctions = urlservice{}

type UrlServiceFunctions interface {
	CreateUrl(dto model.UrlDto) (model.UrlDto, error)
	DeleteUrl(id string) error
	RedirectUrl(shortUrl string) (string, error)
}

type urlservice struct{}

func InitUrlService() {
	urldatastore = datastore.InitUrlDatastore(sqlite3helper.DbConn)
	historydatastore = datastore.InitHistoryDatastore(sqlite3helper.DbConn)

	urlrepo = urlrepository.Init(urldatastore)
	histrepo = historyrepository.Init(historydatastore)
}

func (us urlservice) CreateUrl(dto model.UrlDto) (model.UrlDto, error) {

	err := dto.ValidateCreate()
	if dto.ValidateCreate() != nil {
		return model.UrlDto{}, err
	}

	url, err := urlrepo.Create(context.Background(), datastore.Url{
		LongUrl:      dto.LongUrl,
		ExpirationDt: dto.ExpirationDt,
	})

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			logservice.LogInfo(fmt.Sprintf(messages.SHORT_URL_EXISTS, dto.LongUrl))
			url, err := urlrepo.GetShortUrlByLongUrl(context.Background(), dto.LongUrl)
			if err != nil {
				return model.UrlDto{}, err
			}
			return model.UrlDto{
				UrlID:        url.UrlID,
				LongUrl:      url.LongUrl,
				ShortUrl:     fmt.Sprintf("http://localhost:9000/%s", url.ShortUrl),
				ExpirationDt: url.ExpirationDt,
			}, nil
		} else {
			return model.UrlDto{}, err
		}
	}
	return model.UrlDto{
		UrlID:        url.UrlID,
		LongUrl:      url.LongUrl,
		ShortUrl:     fmt.Sprintf("http://localhost:9000/%s", url.ShortUrl),
		ExpirationDt: url.ExpirationDt,
	}, nil
}

func (us urlservice) DeleteUrl(id string) error {
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	rows, err := histrepo.Delete(context.Background(), int32(convertedId))
	logservice.LogInfo(fmt.Sprintf("%d rows deleted from history", rows))
	return urlrepo.DeleteUrl(context.Background(), int32(convertedId))
}

func (us urlservice) RedirectUrl(shortUrl string) (string, error) {
	url, err := urlrepo.GetLongUrl(context.Background(), shortUrl)
	if err != nil {
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			return "", &helper.CustomError{Message: fmt.Sprintf(messages.SHORT_URL_DOES_NOT_EXIST, shortUrl)}
		}
		return "", err
	}
	histrepo.Insert(context.Background(), url.UrlID)
	return url.LongUrl, nil
}
