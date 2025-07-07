package models

type RoleCreateRequest struct {
	Role string `json:"name" validate:"required"`
}

type RoleUpdateRequest struct {
	ID   int    `json:"id" validate:"required"`
	Role string `json:"name" validate:"required"`
}
