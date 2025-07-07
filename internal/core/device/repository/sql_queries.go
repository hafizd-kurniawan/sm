package repository

const (
	CreateDevice = `
		INSERT INTO devices (name, location, status, created_by) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, location, status;
	`

	GetAllDevices = `
		SELECT id, name, location, status
		FROM devices 
		WHERE deleted_at IS NULL;
	`

	GetDeviceByID = `
		SELECT id, name, location, status
		FROM devices 
		WHERE id = $1 AND deleted_at IS NULL;
	`

	UpdateDevice = `
		UPDATE devices 
		SET name = $1, location = $2, status = $3, updated_by = $4, updated_at = NOW() 
		WHERE id = $5 AND deleted_at IS NULL;
	`

	DeleteDevice = `
		UPDATE devices 
		SET deleted_at = NOW(), deleted_by = $1 
		WHERE id = $2 AND deleted_at IS NULL;
	`
	GetDeviceByName = `
        SELECT id, name, location, status, created_at, updated_at FROM devices WHERE name = $1;
`
)
