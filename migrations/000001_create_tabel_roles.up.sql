CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),           
    updated_at TIMESTAMP,
    updated_by VARCHAR(100),           
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100)
);


INSERT INTO roles (name, created_by)
VALUES 
  ('admin', 'system'),
  ('technician', 'system'),
  ('viewer', 'system');
