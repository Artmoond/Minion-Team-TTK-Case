package models

import "time"

type User struct {
	ID         int64     `json:"id" db:"id"`
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	MiddleName string    `json:"middle_name" db:"middle_name"`
	Roles      []string  `json:"roles" db:"roles"`
	Date       time.Time `json:"date" db:"date"`
}

type GetAllUsersRequest struct {
	Token  string `json:"token"`
	Column string `json:"column"`
	IsASC  bool   `json:"is_asc"`
}

type DeleteUserRequest struct {
	Token      string `json:"token" validate:"required"`
	ToDeleteId int64  `json:"to_delete_id" validate:"required"`
}

type UpdateUserRoleReq struct {
	ID  int64   `json:"id" db:"id"`
	Rol []Roles `json:"roles" db:"roles"`
}

type Roles struct {
	Role     string `json:"role" db:"role"`
	IsAccept bool   `json:"isAccepted" db:"is_accepted"`
}

type ChangePasswordReq struct {
	Token    string `json:"token" validate:"required"`
	ID       int64  `json:"id" db:"id" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}
