package repository

import (
	"context"

	serviceModels "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository/models"
)

type Repository interface {
	GetAllUsers(ctx context.Context, column string, isASC bool) ([]serviceModels.User, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUserRole(ctx context.Context, req *models.UpdateUserRoleReq) error
	UpdatePassword(ctx context.Context, id int64, newPassword []byte) error
}
