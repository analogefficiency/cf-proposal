package historyrepository

import (
	"cf-proposal/common/logservice"
	"cf-proposal/domain/datastore"
	"context"
	"fmt"
	"time"
)

type HistoryInterface interface {
	Create(ctx context.Context, id int32) (datastore.History, error)
}

type History struct {
	repo HistoryInterface
}

func Init(repo HistoryInterface) *History {
	return &History{
		repo: repo,
	}
}

func (h *History) Insert(ctx context.Context, id int32) (datastore.History, error) {
	inserted, err := h.repo.Create(ctx, id)
	if err != nil {
		return datastore.History{}, fmt.Errorf("%w", err)
	}
	logservice.LogInfo(fmt.Sprintf("Writing short url %d accessed %s to history", id, time.Now().Format("2006-01-02 15:04:05")))
	return inserted, nil
}
