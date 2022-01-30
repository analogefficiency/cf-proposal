package api

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/model"
	"cf-proposal/domain/service"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUrlOk(t *testing.T) {

	url := "https://hawaiinewsnow.com"
	service.UrlService = service.UrlServiceMock{}
	service.MockUrlDto = model.UrlDto{
		UrlID:    1,
		LongUrl:  url,
		ShortUrl: fmt.Sprintf("http://localhost:9000/%s", helper.GetShortUrl(url)),
	}

	req, err := http.NewRequest("POST", "/url/create", helper.GetEncodedPayload(model.UrlDto{LongUrl: url}))
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(UrlController{}.HandleCreate)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong status code returned: got %d wanted %d", r.Code, http.StatusOK)
	}
}

func TestCreateUrlErr(t *testing.T) {
	url := "https://hawaiinewsnow.com"
	service.UrlService = service.UrlServiceMock{}
	service.MockUrlDto = model.UrlDto{
		UrlID:    1,
		LongUrl:  url,
		ShortUrl: fmt.Sprintf(helper.GetShortUrl("https://hawaiinewsnow.com")),
	}
	service.MockError = &helper.CustomError{Message: "Serivce error"}

	req, err := http.NewRequest("POST", "/url/create", helper.GetEncodedPayload(model.UrlDto{LongUrl: url}))
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(UrlController{}.HandleCreate)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code returned: got %d wanted %d", r.Code, http.StatusOK)
	}
}
