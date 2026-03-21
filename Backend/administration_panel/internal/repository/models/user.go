package models

import "time"

type User struct {
	ID         int64     `json:"id", db:"id"`
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	MiddleName string    `json:"middle_name" db:"middle_name"`
	Roles      []string  `json:"roles" db:"roles"`
	Date       time.Time `json:"date" db:"date"`
}
