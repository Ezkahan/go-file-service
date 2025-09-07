CREATE TABLE IF NOT EXISTS files (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    icon_path TEXT,
    file_path TEXT NOT NULL,
    category_id UUID REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);