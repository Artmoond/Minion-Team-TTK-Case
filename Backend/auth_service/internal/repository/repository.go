package repository

import (
	"context"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.CreateUserRequest) (*models.CreateUserResponse, error)
}
