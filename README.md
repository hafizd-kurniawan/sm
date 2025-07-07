# Sql Schema
```sql
-- ROLE
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

-- USER
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    role_id INT NOT NULL REFERENCES roles(id),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),           
    updated_at TIMESTAMP,
    updated_by VARCHAR(100),           
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100)

);

INSERT INTO users (
    email, name, password, role_id, created_by
) VALUES (
    'admin@iot.local',
    'Super Admin',
    '$2a$10$Wc3D9qv9hxJhL8u.XOvPlO5B4bwR08K1Hc8F0gBj1cExDf9NLTz3e', -- password: admin123
    1,
    'system'
);

-- DEVICE
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


```
# Migrate
```
 migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/smart_devices?sslmode=disable" up
```

# Start Server
```
❯ go run cmd/api/main.go 
```
