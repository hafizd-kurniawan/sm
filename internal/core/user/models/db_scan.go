package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Role     int    `json:"role" db:"role"`
}
