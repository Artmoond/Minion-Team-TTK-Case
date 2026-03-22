package auth_service

import (
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/repository"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/service"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/service/token"
)

type authService struct {
	repo  repository.Repository
	token token.TokenService
}

func NewAuthService(repo repository.Repository, token token.TokenService) service.Service {
	return &authService{
		repo:  repo,
		token: token,
	}
}
