CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100),
    status VARCHAR(10) NOT NULL CHECK (status IN ('online', 'offline')),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),           
    updated_at TIMESTAMP,
    updated_by VARCHAR(100),           
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100)
);
