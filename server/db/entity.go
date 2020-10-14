package db

import "database/sql"

type Entity struct {
	*Queries
	DB *sql.DB
}

func NewEntity(db *sql.DB) *Entity {
	return &Entity{
		Queries: New(db),
		DB:      db,
	}
}
