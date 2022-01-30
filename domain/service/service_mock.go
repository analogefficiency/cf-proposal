package service

import "cf-proposal/domain/model"

var (
	MockUrlDto      model.UrlDto
	MockShortUrlDto model.LongUrlDto
	MockError       error
	ShortUrl        string
)

type UrlServiceMock struct{}

func (usm UrlServiceMock) CreateUrl(dto model.UrlDto) (model.UrlDto, error) {
	return MockUrlDto, MockError
}

func (usm UrlServiceMock) DeleteUrl(id string) error {
	return MockError
}

func (usm UrlServiceMock) RedirectUrl(shortUrl string) (string, error) {
	return ShortUrl, MockError
}
