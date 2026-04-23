CREATE TABLE games (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    provider TEXT NOT NULL,
    category TEXT NOT NULL,
    rtp REAL NOT NULL,
    variance TEXT NOT NULL CHECK (variance IN ('Low', 'Mid', 'High')),
    enabled BOOLEAN NOT NULL DEFAULT 0,
    launch_date TEXT NOT NULL,
    tags TEXT NOT NULL, -- Will store your array of strings as a JSON string
    image_url TEXT NOT NULL
);
