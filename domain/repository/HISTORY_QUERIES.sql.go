// Code generated by sqlc. DO NOT EDIT.
// source: HISTORY_QUERIES.sql

package repository

import (
	"context"
)

const insertUrlHistory = `-- name: InsertUrlHistory :one
INSERT INTO HISTORY (url_id, access_dt)
VALUES($1, $2) RETURNING history_id, url_id, access_dt
`

type InsertUrlHistoryParams struct {
	UrlID    int32  `json:"url_id"`
	AccessDt string `json:"access_dt"`
}

func (q *Queries) InsertUrlHistory(ctx context.Context, arg InsertUrlHistoryParams) (History, error) {
	row := q.db.QueryRowContext(ctx, insertUrlHistory, arg.UrlID, arg.AccessDt)
	var i History
	err := row.Scan(&i.HistoryID, &i.UrlID, &i.AccessDt)
	return i, err
}
