package db

import (
	"database/sql"
)

// DB wraps a db connection to the companies database
type DB struct {
	db *sql.DB
}

// NewDB returns a new DB
func NewDB(db *sql.DB) *DB {
	return &DB{db: db}
}
