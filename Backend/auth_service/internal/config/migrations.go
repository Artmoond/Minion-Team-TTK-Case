package config

type MigrationsConfig struct {
	migrationPath string
}

func NewMigrationsConfig() *MigrationsConfig {
	return &MigrationsConfig{
		migrationPath: "./migrations",
	}
}

func (c *MigrationsConfig) MigrationPath() string {
	return c.migrationPath
}
