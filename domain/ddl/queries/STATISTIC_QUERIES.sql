-- name: GetStatisticsByUrl :one
SELECT * FROM STATISTICS WHERE url_id = $1;