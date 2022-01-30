package service

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repo/historyrepository"
	"cf-proposal/domain/repo/statisticsrepository"
	"cf-proposal/domain/repo/urlrepository"
	"database/sql"
	"fmt"
	"testing"
)

func TestCreateUrlOk(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	urlrepository.MockCreateUrl = datastore.Url{
		UrlID:        1,
		LongUrl:      url,
		ShortUrl:     helper.GetShortUrl(url),
		ExpirationDt: "",
	}
	urlrepository.MockCreateError = nil

	created, err := UrlService{}.CreateUrl(model.UrlDto{
		LongUrl: url,
	})
	if err != nil {
		t.Fatal("%w", err)
	}

	expected := fmt.Sprintf("http://localhost:9000/%s", helper.GetShortUrl(url))
	if created.ShortUrl != expected {
		t.Errorf(fmt.Sprintf("Short URL expected: %s got %s", created.ShortUrl, expected))
	}
}

func TestCreateUrlOkExisting(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	urlrepository.MockGetShortUrl = datastore.Url{
		UrlID:        1,
		LongUrl:      url,
		ShortUrl:     helper.GetShortUrl(url),
		ExpirationDt: "",
	}
	urlrepository.MockGetUrlError = fmt.Errorf("UNIQUE constraint failed")
	urlrepository.MockGetShortUrlError = nil

	created, err := UrlService{}.CreateUrl(model.UrlDto{
		LongUrl: url,
	})
	if err != nil {
		t.Errorf("Creation error")
	}

	expected := fmt.Sprintf("http://localhost:9000/%s", helper.GetShortUrl(url))
	if created.ShortUrl != expected {
		t.Errorf(fmt.Sprintf("Short URL expected: %s got %s", created.ShortUrl, expected))
	}
}

func TestCreateUrlFail(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	urlrepository.MockGetShortUrl = datastore.Url{
		UrlID:        1,
		LongUrl:      url,
		ShortUrl:     helper.GetShortUrl(url),
		ExpirationDt: "",
	}
	urlrepository.MockCreateError = fmt.Errorf("Random repo error")
	urlrepository.MockGetShortUrlError = nil

	_, err := UrlService{}.CreateUrl(model.UrlDto{
		LongUrl: url,
	})
	if err.Error() != "Random repo error" {
		t.Errorf("Error on repo create not generated")
	}
}

func TestRedirectUrlOK(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	histrepo = historyrepository.Init(historyrepository.HistoryRepoMock{})
	urlrepository.MockGetLongUrl = datastore.Url{
		LongUrl: url,
	}
	urlrepository.MockGetLongUrlError = nil

	longUrl, err := UrlService{}.RedirectUrl(helper.GetShortUrl(url))
	if err != nil {
		t.Errorf("Unexpected error")
	}
	if longUrl != url {
		t.Errorf(fmt.Sprintf("Unexpected URL got %s, expected %s", longUrl, url))
	}
}

func TestRedirectUrlNonExistFail(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	histrepo = historyrepository.Init(historyrepository.HistoryRepoMock{})
	urlrepository.MockGetLongUrl = datastore.Url{}
	urlrepository.MockGetLongUrlError = sql.ErrNoRows

	_, err := UrlService{}.RedirectUrl(helper.GetShortUrl(url))
	if err.Error() != fmt.Sprintf(messages.SHORT_URL_DOES_NOT_EXIST, helper.GetShortUrl(url)) {
		t.Errorf("Unexpected error")
	}
}

func TestRedirectUrlFail(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	histrepo = historyrepository.Init(historyrepository.HistoryRepoMock{})
	urlrepository.MockGetLongUrl = datastore.Url{}
	urlrepository.MockGetLongUrlError = &helper.CustomError{Message: "Unknown repo error"}
	_, err := UrlService{}.RedirectUrl(helper.GetShortUrl(url))
	if err.Error() != "Unknown repo error" {
		t.Errorf("Expected an error, recieved nothing.")
	}
}

func TestDeleteOk(t *testing.T) {
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	histrepo = historyrepository.Init(historyrepository.HistoryRepoMock{})
	historyrepository.MockDeletedRows = 1
	err := UrlService{}.DeleteUrl("1")
	if err != nil {
		t.Errorf("Expected no error.")
	}
}

func TestGetStatisticOk(t *testing.T) {
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	statrepo = statisticsrepository.Init(statisticsrepository.StatisticsRepoMock{})

	statisticsrepository.MockGetStatisticEntity = datastore.Statistic{
		UrlID:           1,
		TwentyFourHours: 1,
		LastSevenDays:   2,
		AllTime:         4,
	}
	statisticsrepository.MockGetStatisticError = nil

	urlrepository.MockGetUrlError = nil

	statistic, err := StatisticsService{}.GetStatistic("1")
	if err != nil {
		t.Errorf("Expected no error.")
	}

	if statisticsrepository.MockGetStatisticEntity != datastore.Statistic(statistic) {
		t.Errorf("Expected statistic not matching struct sent to mock")
	}
}

func TestGetStatisticFail(t *testing.T) {
	urlrepo = urlrepository.Init(urlrepository.UrlRepoMock{})
	statrepo = statisticsrepository.Init(statisticsrepository.StatisticsRepoMock{})

	statisticsrepository.MockGetStatisticEntity = datastore.Statistic{}
	statisticsrepository.MockGetStatisticError = nil
	urlrepository.MockGetUrlError = &helper.CustomError{Message: "Not found"}

	_, err := StatisticsService{}.GetStatistic("1")
	if err.Error() != "Not found" {
		t.Errorf("Expected a not found error")
	}
}
