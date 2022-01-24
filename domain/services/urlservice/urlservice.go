package urlservice

import (
	"cf-proposal/domain/model"
	"context"
	"fmt"
)

type UrlRepo interface {
	Create(ctx context.Context, urlDto model.UrlDto) (model.UrlDto, error)
	GetLongUrl(ctx context.Context, shortUrl string) (model.LongUrlDto, error)
}

type Url struct {
	repo UrlRepo
}

func Init(repo UrlRepo) *Url {
	return &Url{
		repo: repo,
	}
}

func (u *Url) Create(ctx context.Context, urlDto model.UrlDto) (model.UrlDto, error) {
	urlDto, err := u.repo.Create(ctx, urlDto)
	if err != nil {
		return model.UrlDto{}, fmt.Errorf("%w", err)
	}
	return urlDto, nil
}

func (u *Url) GetLongUrl(ctx context.Context, shortUrl string) (model.LongUrlDto, error) {
	longUrl, err := u.repo.GetLongUrl(ctx, shortUrl)
	if err != nil {
		return model.LongUrlDto{}, fmt.Errorf("%w", err)
	}
	return longUrl, nil
}
