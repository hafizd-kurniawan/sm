package repository

const (
	CreateRole = `
		INSERT INTO roles (name, created_by) 
		VALUES ($1, $2)
		RETURNING name;
	`

	GetAllRole = `
		SELECT id, name
		FROM roles
		WHERE deleted_at IS NULL;
	`

	GetRoleByID = `
		SELECT id, name
		FROM roles 
		WHERE id = $1 AND deleted_at IS NULL;
	`

	GetRoleByRole = `
		SELECT id, name FROM roles WHERE name = $1 AND deleted_at IS NULL;
	`

	UpdateRole = `
		UPDATE roles 
		SET name = $1, updated_by = $2, updated_at = NOW() 
		WHERE id = $3 AND deleted_at IS NULL RETURNING name;
	`

	DeleteRole = `
		UPDATE roles 
		SET deleted_at = NOW(), deleted_by = $1 
		WHERE id = $2 AND deleted_at IS NULL;
	`
)
