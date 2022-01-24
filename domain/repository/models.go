// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
)

type History struct {
	HistoryID int32  `json:"history_id"`
	UrlID     int32  `json:"url_id"`
	AccessDt  string `json:"access_dt"`
}

type Url struct {
	UrlID        int32         `json:"url_id"`
	LongUrl      string        `json:"long_url"`
	ShortUrl     string        `json:"short_url"`
	ExpirationDt sql.NullInt32 `json:"expiration_dt"`
}
