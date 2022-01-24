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
		LongUrl:  urlDto.LongUrl,
		ShortUrl: helper.GetShortUrl(urlDto.LongUrl),
		ExpirationDt: sql.NullInt32{
			Int32: urlDto.ExpirationDt,
		},
	})
	if err != nil {
		return model.UrlDto{}, fmt.Errorf("Error: %w", err)
	}
	return model.UrlDto{
		UrlID:        url.UrlID,
		LongUrl:      url.LongUrl,
		ShortUrl:     fmt.Sprintf("http://localhost:9000/%s", helper.GetShortUrl(url.ShortUrl)),
		ExpirationDt: url.ExpirationDt.Int32,
	}, nil
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
