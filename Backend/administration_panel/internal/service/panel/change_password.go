package panel

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (s *Panel) ChangePassword(ctx context.Context, req *models.ChangePasswordReq) error {
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
		}
	}

	if !isAdmin {
		log.Println("Error is not admin")
		return custom_err.ErrNotHaveRightRole
	}

	hashPasswords, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return custom_err.ErrInvalidArguments
	}

	err = s.repo.UpdatePassword(ctx, req.ID, hashPasswords)
	if err != nil {
		return err
	}

	return nil
}
