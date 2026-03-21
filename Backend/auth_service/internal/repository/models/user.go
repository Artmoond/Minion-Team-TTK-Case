package models

import "time"

type UserModel struct {
	Login      string    `json:"login" db:"login"`
	Password   []byte    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	MiddleName string    `json:"middle_name" db:"middle_name"`
	Date       time.Time `json:"date" db:"date"`
}

type CreateUserRequest struct {
	Login      string `json:"login" db:"login"`
	Password   []byte `json:"password" db:"password"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	Roles      []string
}

type CreateUserResponse struct {
	ID    int64    `json:"id"`
	Roles []string `json:"roles"`
}

type FindByNameResponse struct {
	ID       int64    `json:"id"`
	Login    string   `json:"login"`
	Password []byte   `json:"password"`
	Roles    []string `json:"roles"`
}
