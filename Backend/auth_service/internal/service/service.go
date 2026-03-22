package service

import (
	"context"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/models"
)

type Service interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error)
	LoginService(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
}
