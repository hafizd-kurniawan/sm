package models

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserCreateResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserListResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
