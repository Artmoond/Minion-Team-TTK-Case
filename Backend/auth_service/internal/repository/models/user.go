package models

import "time"

type UserModel struct {
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	MiddleName string    `json:"middle_name" db:"middle_name"`
	Role       string    `json:"role" db:"role"`
	Date       time.Time `json:"date" db:"date"`
}

type CreateUserRequest struct {
	Login      string `json:"login" db:"login"`
	Password   string `json:"password" db:"password"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	Role       string `json:"role" db:"role"`
}

type CreateUserResponse struct {
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	MiddleName string    `json:"middle_name" db:"middle_name"`
	Role       string    `json:"role" db:"role"`
	Date       time.Time `json:"date" db:"date"`
}
