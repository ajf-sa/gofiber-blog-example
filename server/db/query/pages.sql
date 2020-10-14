-- name: CreatePage :one
INSERT INTO pages(title,slug,body)values($1,$2,$3) RETURNING * ;

-- name: ListPage :many
SELECT * FROM pages ;


-- name: GetPageBySlug :one
select * from pages where slug=$1 limit 1 offset 0;