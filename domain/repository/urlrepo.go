package repository

import (
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
		UrlID:    urlDto.UrlID,
		LongUrl:  urlDto.LongUrl,
		ShortUrl: urlDto.ShortUrl,
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
		ShortUrl:     url.ShortUrl,
		ExpirationDt: url.ExpirationDt.Int32,
	}, nil
}
