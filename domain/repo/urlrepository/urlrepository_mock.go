package urlrepository

import (
	"cf-proposal/domain/datastore"
	"context"
)

var (
	MockCreateUrl        datastore.Url
	MockGetUrl           datastore.Url
	MockGetLongUrl       datastore.Url
	MockGetShortUrl      datastore.Url
	MockError            error
	MockGetShortUrlError error
	MockGetLongUrlError  error
	MockCreateError      error
	MockDeleteError      error
)

type UrlRepoMock struct{}

func (urm UrlRepoMock) Create(ctx context.Context, urlDto datastore.Url) (datastore.Url, error) {
	return MockCreateUrl, MockCreateError
}

func (urm UrlRepoMock) DeleteUrl(ctx context.Context, id int32) error {
	return MockDeleteError
}

func (urm UrlRepoMock) GetUrl(ctx context.Context, id int32) (datastore.Url, error) {
	return MockGetUrl, MockError
}

func (urm UrlRepoMock) GetLongUrl(ctx context.Context, shortUrl string) (datastore.Url, error) {
	return MockGetLongUrl, MockGetLongUrlError
}

func (urm UrlRepoMock) GetShortUrlByLongUrl(ctx context.Context, longUrl string) (datastore.Url, error) {
	return MockGetShortUrl, MockGetShortUrlError
}
