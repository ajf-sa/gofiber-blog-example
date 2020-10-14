// Code generated by sqlc. DO NOT EDIT.
// source: pages.sql

package db

import (
	"context"
)

const createPage = `-- name: CreatePage :one
INSERT INTO pages(title,slug,body)values($1,$2,$3) RETURNING id, title, slug, body, is_active, create_at
`

type CreatePageParams struct {
	Title string
	Slug  string
	Body  string
}

func (q *Queries) CreatePage(ctx context.Context, arg CreatePageParams) (Page, error) {
	row := q.queryRow(ctx, q.createPageStmt, createPage, arg.Title, arg.Slug, arg.Body)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Body,
		&i.IsActive,
		&i.CreateAt,
	)
	return i, err
}

const getPageBySlug = `-- name: GetPageBySlug :one
select id, title, slug, body, is_active, create_at from pages where slug=$1 limit 1 offset 0
`

func (q *Queries) GetPageBySlug(ctx context.Context, slug string) (Page, error) {
	row := q.queryRow(ctx, q.getPageBySlugStmt, getPageBySlug, slug)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Body,
		&i.IsActive,
		&i.CreateAt,
	)
	return i, err
}

const listPage = `-- name: ListPage :many
SELECT id, title, slug, body, is_active, create_at FROM pages
`

func (q *Queries) ListPage(ctx context.Context) ([]Page, error) {
	rows, err := q.query(ctx, q.listPageStmt, listPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Page
	for rows.Next() {
		var i Page
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Slug,
			&i.Body,
			&i.IsActive,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
