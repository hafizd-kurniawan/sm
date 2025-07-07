package exception

import "errors"

// Standard application errors
var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("data already exists")
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
)
