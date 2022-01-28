-- name: CreateUrl :one
INSERT INTO URL (long_url, short_url, expiration_dt)
VALUES($1, $2, $3) RETURNING *;

-- name: DeleteUrl :exec
DELETE FROM URL 
WHERE url_id = $1;

-- name: FindRedirectByShortUrl :one
SELECT url_id, long_url FROM URL
WHERE short_url = $1;

-- name: FindShortUrlByLongUrl :one
SELECT * FROM URL
WHERE long_url = $1;