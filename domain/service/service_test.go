package service

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/datastore"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repo/historyrepository"
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
	urlrepository.MockError = fmt.Errorf("UNIQUE constraint failed")
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
