package urlservice

import (
	"cf-proposal/domain/model"
	"context"
	"fmt"
)

type UrlRepo interface {
	Create(ctx context.Context, urlDto model.UrlDto) (model.UrlDto, error)
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
		return model.UrlDto{}, fmt.Errorf("error: %w", err)
	}
	return urlDto, nil
}
