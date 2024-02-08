CREATE TABLE IF NOT EXISTS books
(
    id      INTEGER PRIMARY KEY,
    author  TEXT NOT NULL,
    title   TEXT NOT NULL,
    country TEXT NOT NULL
);