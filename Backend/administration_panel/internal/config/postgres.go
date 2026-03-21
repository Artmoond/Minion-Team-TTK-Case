package config

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	host         string
	port         string
	databaseName string
	user         string
	password     string
}

func NewPostgresConfig() *PostgresConfig {
	host := os.Getenv("PG_HOST")
	if host == "" {
		host = "localhost"
	}

	return &PostgresConfig{
		host:         host,
		port:         os.Getenv("PG_PORT"),
		databaseName: os.Getenv("PG_DATABASE_NAME"),
		user:         os.Getenv("PG_USER"),
		password:     os.Getenv("PG_PASSWORD"),
	}
}

func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.host,
		c.port,
		c.databaseName,
		c.user,
		c.password,
	)
}
