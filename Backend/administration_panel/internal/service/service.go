package service

import (
	"context"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
)

type PanelService interface {
	GetAllUsers(ctx context.Context, req *models.GetAllUsersRequest) ([]models.User, error)
	DeleteUser(ctx context.Context, req *models.DeleteUserRequest) error
	UpdateUserRole(ctx context.Context, token string, req *models.UpdateUserRoleReq) error
}
