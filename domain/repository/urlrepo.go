package repository

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/model"
	"context"
	"database/sql"
	"fmt"
)

type UrlRepo struct {
	q *Queries
}

func InitUrlRepo(db *sql.DB) *UrlRepo {
	return &UrlRepo{
		q: New(db),
	}
}

func (u *UrlRepo) Create(ctx context.Context, urlDto model.UrlDto) (model.UrlDto, error) {
	url, err := u.q.CreateUrl(ctx, CreateUrlParams{
		LongUrl:      urlDto.LongUrl,
		ShortUrl:     helper.GetShortUrl(urlDto.LongUrl),
		ExpirationDt: urlDto.ExpirationDt,
	})
	if err != nil {
		return model.UrlDto{}, fmt.Errorf("%w", err)
	}
	return model.UrlDto{
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

func (u *UrlRepo) GetLongUrl(ctx context.Context, shortUrl string) (model.LongUrlDto, error) {
	returnValue, err := u.q.FindRedirectByShortUrl(ctx, shortUrl)
	if err != nil {
		return model.LongUrlDto{}, fmt.Errorf("%w", err)
	}
	return model.LongUrlDto{
		UrlID:   returnValue.UrlID,
		LongUrl: returnValue.LongUrl,
	}, nil
}

func (u *UrlRepo) GetShortUrlByLongUrl(ctx context.Context, longUrl string) (model.UrlDto, error) {
	url, err := u.q.FindShortUrlByLongUrl(ctx, longUrl)
	if err != nil {
		return model.UrlDto{}, fmt.Errorf("%w", err)
	}
	return model.UrlDto{
		UrlID:        url.UrlID,
		LongUrl:      url.LongUrl,
		ShortUrl:     fmt.Sprintf("http://localhost:9000/%s", url.ShortUrl),
		ExpirationDt: url.ExpirationDt,
	}, nil
}
