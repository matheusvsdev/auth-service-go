CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT,
    provider TEXT NOT NULL, --local | google | github
    plan TEXT NOT NULL, --basic | premium
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);