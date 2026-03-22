package panel

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
)

const RoleAdmin = "Администратор"

func (s *Panel) GetAllUsers(ctx context.Context, req *models.GetAllUsersRequest) ([]models.User, error) {
	if len(req.Token) == 0 {
		log.Println("err token is empty: ", req.Token)
		return nil, custom_err.ErrEmptyToken
	}

	claims, err := s.token.GetClaims(req.Token)
	if err != nil {
		return nil, err
	}

	isAdmin := false

	for _, v := range claims.Roles {
		if v == RoleAdmin {
			isAdmin = true
			break
		}

		isAdmin = false
	}

	if !isAdmin {
		log.Println("err don`t have right role")
		return nil, custom_err.ErrNotHaveRightRole
	}

	resp, err := s.repo.GetAllUsers(ctx, req.Column, req.IsASC)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
