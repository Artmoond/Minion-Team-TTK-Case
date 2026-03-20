package migrator

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

type Migrator struct {
	db             *sql.DB
	migrationsPath string
}

func NewMigrator(db *sql.DB, migrationsPath string) *Migrator {
	return &Migrator{
		db:             db,
		migrationsPath: migrationsPath,
	}
}

func (m *Migrator) Up() error {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Print("custom_err set migration dialect: ", err)
		return err
	}

	err := goose.Up(m.db, m.migrationsPath)
	if err != nil {
		log.Print("custom_err migration up: ", err)
		return err
	}

	return nil
}
