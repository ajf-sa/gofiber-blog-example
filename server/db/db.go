// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createPageStmt, err = db.PrepareContext(ctx, createPage); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePage: %w", err)
	}
	if q.getPageBySlugStmt, err = db.PrepareContext(ctx, getPageBySlug); err != nil {
		return nil, fmt.Errorf("error preparing query GetPageBySlug: %w", err)
	}
	if q.listPageStmt, err = db.PrepareContext(ctx, listPage); err != nil {
		return nil, fmt.Errorf("error preparing query ListPage: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createPageStmt != nil {
		if cerr := q.createPageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPageStmt: %w", cerr)
		}
	}
	if q.getPageBySlugStmt != nil {
		if cerr := q.getPageBySlugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPageBySlugStmt: %w", cerr)
		}
	}
	if q.listPageStmt != nil {
		if cerr := q.listPageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPageStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                DBTX
	tx                *sql.Tx
	createPageStmt    *sql.Stmt
	getPageBySlugStmt *sql.Stmt
	listPageStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                tx,
		tx:                tx,
		createPageStmt:    q.createPageStmt,
		getPageBySlugStmt: q.getPageBySlugStmt,
		listPageStmt:      q.listPageStmt,
	}
}
