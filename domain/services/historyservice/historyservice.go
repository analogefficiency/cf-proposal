package historyservice

import (
	"cf-proposal/common/logservice"
	"cf-proposal/domain/repository"
	"context"
	"fmt"
	"time"
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
	logservice.LogInfo(fmt.Sprintf("Writing short url %d accessed %s to history", id, time.Now().Format("2006-01-02 15:04:05")))
	return inserted, nil
}
