package repository

const (
	CreateUser = `
		INSERT INTO users (name, email, password, role_id, created_by) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING name, email;
	`

	GetAllUsers = `
		SELECT u.id, u.name, u.email, r.name as role_name
		FROM users u
		JOIN roles r ON u.role_id = r.id
		WHERE u.deleted_at IS NULL;
	`

	GetUserByID = `
		SELECT id, name, email, password, role_id
		FROM users 
		WHERE id = $1 AND deleted_at IS NULL;
	`

	GetUserByEmail = `
		SELECT id, name, email, password, role_id
		FROM users 
		WHERE email = $1 AND deleted_at IS NULL;
	`

	UpdateUser = `
		UPDATE users 
		SET name = $1, email = $2, password = $3, role_id = $4, updated_by = $5, updated_at = NOW() 
		WHERE id = $6 AND deleted_at IS NULL;
	`

	DeleteUser = `
		UPDATE users 
		SET deleted_at = NOW(), deleted_by = $1 
		WHERE id = $2 AND deleted_at IS NULL;
	`
	GetUserByEmailAndRole = `
		SELECT u.id, u.name, u.email, r.name
		FROM users u
		INNER JOIN roles r ON u.role_id = r.id
		WHERE u.email = $1;
    `
)
