-- name: FindRedirectByShortUrl :one
SELECT * FROM URL
WHERE short_url = $1;