CREATE TABLE IF NOT EXISTS plugins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT NOT NULL,
    type_id INTEGER NOT NULL,
    type_name TEXT NOT NULL,
    type_code TEXT NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT NOT NULL,
    org_name TEXT NOT NULL,
    org_url TEXT NOT NULL,
    keywords TEXT NOT NULL,
    version TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    readme TEXT NOT NULL
);
