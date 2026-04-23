-- name: GetIngestionLog :one
-- Retrieves a specific log to check if it has already succeeded.
SELECT id, source_url, status, ingested_at
FROM ingestion_logs
WHERE source_url = ? LIMIT 1;

-- name: StartIngestion :one
-- Inserts a new PENDING log, or resets an existing one to PENDING if it previously FAILED.
INSERT INTO ingestion_logs (source_url, status)
VALUES (?, 'PENDING')
ON CONFLICT (source_url) 
DO UPDATE SET 
    status = EXCLUDED.status,
    ingested_at = CURRENT_TIMESTAMP
RETURNING id, source_url, status, ingested_at;

-- name: UpdateIngestionStatus :exec
-- Marks the ingestion as SUCCESS or FAILED once the worker finishes.
UPDATE ingestion_logs
SET 
    status = ?2, 
    ingested_at = CURRENT_TIMESTAMP
WHERE source_url = ?1;
