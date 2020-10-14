-- name: CreatePage :one
INSERT INTO pages(title,slug,body)values($1,$2,$3) RETURNING * ;