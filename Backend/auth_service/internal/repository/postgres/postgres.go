package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

/*func NewPostgres(db *pgxpool.Pool) repository.Repository {
	return &postgres{db: db}
}*/
