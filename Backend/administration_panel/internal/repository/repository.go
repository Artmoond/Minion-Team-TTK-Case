package repository

import (
	"context"

	serviceModels "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
)

type Repository interface {
	GetAllUsers(ctx context.Context, column string, isASC bool) ([]serviceModels.User, error)
}
