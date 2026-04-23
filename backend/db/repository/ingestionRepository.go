package db

import (
	sqlc "backend/db/sqlc_generated"
	"context"
	"database/sql"
)

type IngestionRepository interface {
	// GetIngestionLog retrieves a specific log to check if it has already succeeded.
	GetIngestionLog(ctx context.Context, sourceUrl string) (sqlc.IngestionLog, error)

	// StartIngestion inserts a new PENDING log, or resets an existing one to PENDING.
	StartIngestion(ctx context.Context, sourceUrl string) (sqlc.IngestionLog, error)

	// UpdateIngestionStatus marks the ingestion as SUCCESS or FAILED.
	UpdateIngestionStatus(ctx context.Context, arg sqlc.UpdateIngestionStatusParams) error
}

type SQLiteIngestionRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewSQLiteIngestionRepository(db *sql.DB) *SQLiteIngestionRepository {
	return &SQLiteIngestionRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

var _ IngestionRepository = (*SQLiteIngestionRepository)(nil)

func (r *SQLiteIngestionRepository) GetIngestionLog(ctx context.Context, sourceUrl string) (sqlc.IngestionLog, error) {
	return r.queries.GetIngestionLog(ctx, sourceUrl)
}

func (r *SQLiteIngestionRepository) StartIngestion(ctx context.Context, sourceUrl string) (sqlc.IngestionLog, error) {
	return r.queries.StartIngestion(ctx, sourceUrl)
}

func (r *SQLiteIngestionRepository) UpdateIngestionStatus(ctx context.Context, arg sqlc.UpdateIngestionStatusParams) error {
	return r.queries.UpdateIngestionStatus(ctx, arg)
}
