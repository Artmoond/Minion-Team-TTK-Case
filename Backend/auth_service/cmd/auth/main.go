package main

import (
	"context"
	"log"
	"time"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/config"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/handlers"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/migrator"
	postgres "github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/postgres"
	authservice "github.com/Artmoond/Minion-Team-TTK-Case/internal/service/auth_service"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/service/token"
	"github.com/Artmoond/Minion-Team-TTK-Case/lib/validator"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	validator.Init()

	tokenSecret, err := config.NewTokenSecret()
	if err != nil {
		log.Fatal(err)
	}

	postgresConfig := config.NewPostgresConfig()

	db, err := pgxpool.New(ctx, postgresConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(ctx); err != nil {
		log.Fatal("err ping database: ", err)
	}

	sqlDB := stdlib.OpenDB(*db.Config().ConnConfig)
	defer sqlDB.Close()

	migrationConfig := config.NewMigrationsConfig()
	migratorRunner := migrator.NewMigrator(sqlDB, migrationConfig.MigrationPath())

	if err = migratorRunner.Up(); err != nil {
		log.Fatal("err migrator up: ", err)
	}

	repo := postgres.NewPostgres(db)
	tokenService := token.NewTokenService(tokenSecret.Secret(), 24*time.Hour)
	authService := authservice.NewAuthService(repo, tokenService)
	authHandlers := handlers.NewHandlers(authService)

	router := gin.Default()
	authGroup := router.Group("/api/v1/auth")
	authGroup.POST("/register", authHandlers.CreateUser)

	log.Println("starting auth server on :3030")

	if err = router.Run(":3030"); err != nil {
		log.Fatal(err)
	}
}
