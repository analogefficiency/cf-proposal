-- name: InsertUrlHistory :one
INSERT INTO HISTORY (url_id, access_dt)
VALUES($1, $2) RETURNING *;