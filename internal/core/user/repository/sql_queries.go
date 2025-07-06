package repository

const (
	CreateUser = `
	INSERT INTO users (name, email, password, role_id) VALUES ($1, $2, $3, $4) RETURNING name, email;
	`

	GetUserByID = `
		SELECT id, name, email, password, role_id FROM users WHERE id = $1;
	`

	GetUserByEmail = `
		SELECT id, name, email, password, role_id FROM users WHERE email = $1;
	`

	UpdateUser = `
		UPDATE users SET name = $1, email = $2, password = $3, role_id = $4 WHERE id = $5;
	`

	DeleteUser = `
		DELETE FROM users WHERE id = $1;
	`

	DeleteUserByEmail = `
		DELETE FROM users WHERE email = $1;
	`

	GetAllUsers = `
		SELECT id, name, email, password, role_id FROM users;
	`

	Login = `
		SELECT id, name, email, password, role_id FROM users WHERE email = $1 AND password = $2;
	`
)
