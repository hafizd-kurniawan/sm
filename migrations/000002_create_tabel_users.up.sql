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
