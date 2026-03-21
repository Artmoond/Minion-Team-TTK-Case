package repository

import (
	"context"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.CreateUserRequest) (*models.CreateUserResponse, error)
	FindByName(ctx context.Context, login string) (*models.FindByNameResponse, error)
}
