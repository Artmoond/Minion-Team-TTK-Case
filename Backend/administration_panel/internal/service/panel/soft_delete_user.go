package panel

import (
	"context"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	"github.com/go-playground/validator/v10"
)

func (s *Panel) DeleteUser(ctx context.Context, req *models.DeleteUserRequest) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}

	claims, err := s.token.GetClaims(req.Token)
	if err != nil {
		return err
	}

	isAdmin := false

	for _, v := range claims.Roles {
		if v == RoleAdmin {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return custom_err.ErrNotHaveRightRole
	}

	err = s.repo.DeleteUser(ctx, req.ToDeleteId)
	if err != nil {
		return err
	}

	return nil
}
