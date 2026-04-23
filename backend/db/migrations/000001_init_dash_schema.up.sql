CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    provider TEXT NOT NULL,
    category TEXT NOT NULL,
    rtp NUMERIC NOT NULL,
    variance TEXT NOT NULL, -- Storing as TEXT, though you could use an ENUM
    enabled BOOLEAN NOT NULL DEFAULT false,
    launch_date DATE NOT NULL,
    tags TEXT[] NOT NULL DEFAULT '{}',
    image_url TEXT NOT NULL
);
