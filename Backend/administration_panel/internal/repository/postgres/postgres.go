package postgres

import (
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName          = "users"
	rolesTableName     = "roles"
	userRolesTableName = "user_roles"
	idColumn           = "id"
	userIDColumn       = "user_id"
	roleIDColumn       = "role_id"
	loginColumn        = "login"
	passwordColumn     = "password"
	firstNameColumn    = "first_name"
	lastNameColumn     = "last_name"
	middleNameColumn   = "middle_name"
	roleNameColumn     = "rolename"
	dateColumn         = "date"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(db *pgxpool.Pool) repository.Repository {
	return &Postgres{
		db: db,
	}
}
