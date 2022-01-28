package datastore

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"context"
	"database/sql"
	"fmt"
)

type UrlRepo struct {
	q *Queries
}

func InitUrlDatastore(db *sql.DB) *UrlRepo {
	return &UrlRepo{
		q: New(db),
	}
}

func (u *UrlRepo) Create(ctx context.Context, urlDto Url) (Url, error) {
	url, err := u.q.CreateUrl(ctx, CreateUrlParams{
		LongUrl:      urlDto.LongUrl,
		ShortUrl:     helper.GetShortUrl(urlDto.LongUrl),
		ExpirationDt: urlDto.ExpirationDt,
	})
	if err != nil {
		return Url{}, fmt.Errorf("%w", err)
	}
	return Url{
		UrlID:        url.UrlID,
		LongUrl:      url.LongUrl,
		ShortUrl:     fmt.Sprintf("http://localhost:9000/%s", url.ShortUrl),
		ExpirationDt: url.ExpirationDt,
	}, nil
}

func (u *UrlRepo) DeleteUrl(ctx context.Context, id int32) error {
	err := u.q.DeleteUrl(ctx, id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (u *UrlRepo) GetUrl(ctx context.Context, id int32) (Url, error) {
	url, err := u.q.GetUrl(ctx, id)
	if err != nil {
		return Url{}, &helper.CustomError{Message: fmt.Sprintf(messages.ENTITY_DOES_NOT_EXIST, "Short Url", id)}
	}
	return Url{
		UrlID:        url.UrlID,
		LongUrl:      url.LongUrl,
		ShortUrl:     url.ShortUrl,
		ExpirationDt: url.ExpirationDt,
	}, nil
}

func (u *UrlRepo) GetLongUrl(ctx context.Context, shortUrl string) (Url, error) {
	returnValue, err := u.q.FindRedirectByShortUrl(ctx, shortUrl)
	if err != nil {
		return Url{}, fmt.Errorf("%w", err)
	}
	return Url{
		UrlID:   returnValue.UrlID,
		LongUrl: returnValue.LongUrl,
	}, nil
}

func (u *UrlRepo) GetShortUrlByLongUrl(ctx context.Context, longUrl string) (Url, error) {
	url, err := u.q.FindShortUrlByLongUrl(ctx, longUrl)
	if err != nil {
		return Url{}, fmt.Errorf("%w", err)
	}
	return Url{
		UrlID:        url.UrlID,
		LongUrl:      url.LongUrl,
		ShortUrl:     fmt.Sprintf("http://localhost:9000/%s", url.ShortUrl),
		ExpirationDt: url.ExpirationDt,
	}, nil
}
