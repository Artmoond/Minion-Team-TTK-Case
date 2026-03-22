package panel

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/converter"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
)

func (s *Panel) UpdateUserRole(ctx context.Context, token string, req *models.UpdateUserRoleReq) error {

	claims, err := s.token.GetClaims(token)
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
		log.Println("Admin role not found in token", claims.Roles)
		return custom_err.ErrNotHaveRightRole
	}

	err = s.repo.UpdateUserRole(ctx, converter.ToRepoFromService(req))
	if err != nil {
		return err
	}

	return nil
}
