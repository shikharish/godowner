PRAGMA foreign_keys = ON;

CREATE TABLE godowns (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    parent_godown TEXT,
    FOREIGN KEY (parent_godown) REFERENCES godowns (id)
);

CREATE TABLE items (
    item_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    category TEXT,
    price REAL,
    status TEXT,
    godown_id TEXT,
    brand TEXT,
    attributes TEXT,
    image_url TEXT,
    FOREIGN KEY (godown_id) REFERENCES godowns (id) ON DELETE CASCADE
);
