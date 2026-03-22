package panel

import (
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/service/token"
)

type Panel struct {
	repo  repository.Repository
	token token.TokenService
}

func NewPanel(repo repository.Repository, token token.TokenService) *Panel {
	return &Panel{repo, token}
}
