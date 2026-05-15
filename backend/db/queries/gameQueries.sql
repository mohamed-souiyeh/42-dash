-- name: UpsertGame :one
-- Inserts a new game or updates an existing one based on the ID.
INSERT INTO games (
    id, name, provider, category, rtp, variance, enabled, launch_date, tags, image_url
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    provider = EXCLUDED.provider,
    category = EXCLUDED.category,
    rtp = EXCLUDED.rtp,
    variance = EXCLUDED.variance,
    enabled = EXCLUDED.enabled,
    launch_date = EXCLUDED.launch_date,
    tags = EXCLUDED.tags,
    image_url = EXCLUDED.image_url
RETURNING id, name, provider, category, rtp, variance, enabled, launch_date, tags, image_url;


-- name: GetGameByID :one
SELECT *
FROM games
WHERE id = ?;
