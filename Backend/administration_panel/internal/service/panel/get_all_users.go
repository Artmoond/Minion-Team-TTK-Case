package panel

import (
	"context"
	"fmt"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
)

const RoleAdmin = "Администратор"

func (s *Panel) GetAllUsers(ctx context.Context, tokenString string) ([]models.User, error) {
	if len(tokenString) == 0 {
		log.Println("err token is empty: ", tokenString)
		return nil, custom_err.ErrEmptyToken
	}

	claims, err := s.token.GetClaims(tokenString)
	if err != nil {
		return nil, err
	}

	isAdmin := false

	for _, v := range claims.Roles {
		if v == RoleAdmin {
			isAdmin = true
			break
		}

		fmt.Println(v)

		isAdmin = false
	}

	if !isAdmin {
		log.Println("err don`t have right role")
		return nil, custom_err.ErrNotHaveRightRole
	}

	resp, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
