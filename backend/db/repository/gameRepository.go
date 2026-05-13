package db

import (
	sqlc "backend/db/sqlc_generated"
	"context"
	"database/sql"
)

type GameRepository interface {
	UpsertGame(ctx context.Context, arg sqlc.UpsertGameParams) (sqlc.Game, error)
}

type SQLiteGameRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewSQLiteGameRepository(db *sql.DB) *SQLiteGameRepository {
	return &SQLiteGameRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

var _ GameRepository = (*SQLiteGameRepository)(nil)

func (r *SQLiteGameRepository) UpsertGame(ctx context.Context, arg sqlc.UpsertGameParams) (sqlc.Game, error) {
	return r.queries.UpsertGame(ctx, arg)
}
