-- name: FindRedirectByShortUrl :one
SELECT * FROM URL
WHERE short_url = $1;

-- name: CreateUrl :one
INSERT INTO URL (url_id, long_url, short_url, expiration_dt)
VALUES($1, $2, $3, $4) RETURNING *;