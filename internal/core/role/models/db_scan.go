package models

type Role struct {
	ID   string `json:"id" db:"id"`
	Role string `json:"role" db:"role"`
}

