package auth_service

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/models"
	"github.com/Artmoond/Minion-Team-TTK-Case/lib/validator"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) LoginService(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	if err := validator.Validate.Struct(req); err != nil {
		log.Println("Err validation on login: ", err)
		return nil, custom_err.ErrInvalidArguments
	}

	resp, err := s.repo.FindByName(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(resp.Password, []byte(req.Password))
	if err != nil {
		log.Println("Err comparing password: ", err)
		return nil, custom_err.ErrPasswordNotAddUp
	}

	token, err := s.token.GenerateToken(resp.ID, resp.Roles)
	if err != nil {
		log.Println("Err generating token: ", err)
		return nil, custom_err.ErrLoginUser
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}
