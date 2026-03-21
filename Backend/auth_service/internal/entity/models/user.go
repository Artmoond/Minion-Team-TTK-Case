package models

type CreateUserRequest struct {
	Login      string `json:"login" validate:"required,latin"`
	Password   string `json:"password" validate:"required"`
	FirstName  string `json:"first_name" validate:"required,cyrillic"`
	LastName   string `json:"last_name" validate:"required,cyrillic"`
	MiddleName string `json:"middle_name" validate:"required,cyrillic"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
