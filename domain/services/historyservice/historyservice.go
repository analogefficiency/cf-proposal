package historyservice

import (
	"cf-proposal/domain/repository"
	"context"
	"fmt"
)

type HistoryRepo interface {
	Create(ctx context.Context, id int32) (repository.History, error)
}

type History struct {
	repo HistoryRepo
}

func Init(repo HistoryRepo) *History {
	return &History{
		repo: repo,
	}
}

func (h *History) Insert(ctx context.Context, id int32) (repository.History, error) {
	inserted, err := h.repo.Create(ctx, id)
	if err != nil {
		return repository.History{}, fmt.Errorf("%w", err)
	}
	return inserted, nil
}
