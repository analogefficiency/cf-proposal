package datastore

import (
	"cf-proposal/common"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type HistoryDatastore struct {
	q *Queries
}

func InitHistoryDatastore(db *sql.DB) *HistoryDatastore {
	return &HistoryDatastore{
		q: New(db),
	}
}

func (u *HistoryDatastore) Create(ctx context.Context, id int32) (History, error) {
	entry, err := u.q.InsertUrlHistory(ctx, InsertUrlHistoryParams{
		UrlID:    id,
		AccessDt: time.Now().Format(common.DATETIME_FORMAT),
	})
	if err != nil {
		return History{}, fmt.Errorf("%w", err)
	}
	return History{
		UrlID:    entry.UrlID,
		AccessDt: entry.AccessDt,
	}, nil
}
