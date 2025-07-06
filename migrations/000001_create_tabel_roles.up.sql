CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);
