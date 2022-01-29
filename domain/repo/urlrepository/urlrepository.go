package urlrepository

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/datastore"
	"context"
	"fmt"
	"strconv"
)

type UrlInterface interface {
	Create(ctx context.Context, urlDto datastore.Url) (datastore.Url, error)
	DeleteUrl(ctx context.Context, id int32) error
	GetUrl(ctx context.Context, id int32) (datastore.Url, error)
	GetLongUrl(ctx context.Context, shortUrl string) (datastore.Url, error)
	GetShortUrlByLongUrl(ctx context.Context, longUrl string) (datastore.Url, error)
}

type Url struct {
	repo UrlInterface
}

func Init(repo UrlInterface) *Url {
	return &Url{
		repo: repo,
	}
}

func (u *Url) Create(ctx context.Context, urlDto datastore.Url) (datastore.Url, error) {
	urlDto, err := u.repo.Create(ctx, urlDto)
	if err != nil {
		return datastore.Url{}, fmt.Errorf("%w", err)
	}
	return urlDto, nil
}

func (u *Url) DeleteUrl(ctx context.Context, id int32) error {

	err := u.repo.DeleteUrl(ctx, id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (u *Url) GetUrl(ctx context.Context, id string) (datastore.Url, error) {

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return datastore.Url{}, &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	url, err := u.repo.GetUrl(ctx, int32(convertedId))
	if err != nil {
		return datastore.Url{}, err
	}
	return url, nil
}

func (u *Url) GetLongUrl(ctx context.Context, shortUrl string) (datastore.Url, error) {
	longUrl, err := u.repo.GetLongUrl(ctx, shortUrl)
	if err != nil {
		return datastore.Url{}, fmt.Errorf("%w", err)
	}
	return longUrl, nil
}

func (u *Url) GetShortUrlByLongUrl(ctx context.Context, longUrl string) (datastore.Url, error) {
	urlDto, err := u.repo.GetShortUrlByLongUrl(ctx, longUrl)
	if err != nil {
		return datastore.Url{}, fmt.Errorf("%w", err)
	}
	return urlDto, nil
}
