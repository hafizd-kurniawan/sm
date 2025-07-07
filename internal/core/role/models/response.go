package models

type RoleListResponse struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type RoleCreateResponse struct {
	ID   string `json:"id"`
	Role string `json:"name"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}
