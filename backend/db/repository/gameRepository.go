package db

import (
	"database/sql"
	sqlc "echo/db/sqlc_generated"
)

type GameRepository interface {
	// GetUserById(ctx, id) // Read
	// SearchUsersByUsername(ctx, username) // Read - implement it using the LIKE keyword in the query.
}

type PostgresGameRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewPostgresGameRepository(db *sql.DB) *PostgresGameRepository {
	return &PostgresGameRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

var _ GameRepository = (*PostgresGameRepository)(nil)
