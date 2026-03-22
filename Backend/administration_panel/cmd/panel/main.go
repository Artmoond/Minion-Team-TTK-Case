package main

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/config"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/handlers"
	postgres "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository/postgres"
	panelservice "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/service/panel"
	tokenservice "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/service/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

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

	repo := postgres.NewPostgres(db)
	tokenSvc := tokenservice.NewTokenService(tokenSecret.Secret())
	panelSvc := panelservice.NewPanel(repo, tokenSvc)
	panelHandlers := handlers.NewHandlers(panelSvc)

	router := gin.Default()
	adminGroup := router.Group("/api/v1/admin")
	adminGroup.GET("/users/all", panelHandlers.GetAllUsers)
	adminGroup.DELETE("/users/:id", panelHandlers.DeleteUser)
	adminGroup.PUT("/users/change/roles", panelHandlers.UpdateUserRole)
	adminGroup.PATCH("/users/change/password", panelHandlers.UpdatePassword)
	
	log.Println("starting administration panel on :3031")

	if err = router.Run(":3031"); err != nil {
		log.Fatal(err)
	}
}
