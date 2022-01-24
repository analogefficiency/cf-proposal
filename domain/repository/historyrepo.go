package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type HistoryRepo struct {
	q *Queries
}

func InitHistoryRepo(db *sql.DB) *HistoryRepo {
	return &HistoryRepo{
		q: New(db),
	}
}

func (u *HistoryRepo) Create(ctx context.Context, id int32) (History, error) {
	entry, err := u.q.InsertUrlHistory(ctx, InsertUrlHistoryParams{
		UrlID:    id,
		AccessDt: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return History{}, fmt.Errorf("%w", err)
	}
	return History{
		UrlID:    entry.UrlID,
		AccessDt: entry.AccessDt,
	}, nil
}
