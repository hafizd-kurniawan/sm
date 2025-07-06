package repository

const (
	CreateRole = `
	INSERT INTO roles (name) VALUES ($1) RETURNING name;
	`

	GetRoleByID = `
		SELECT id, name FROM roles WHERE id = $1;
	`

	GetRoleByRole = `
		SELECT id, name FROM roles WHERE role = $1;
	`

	UpdateRole = `
		UPDATE roles SET name = $1 WHERE id = $2;
	`

	DeleteRole = `
		DELETE FROM roles WHERE id = $1;
	`

	GetAllRole = `
		SELECT id, name FROM roles;
	`
)
