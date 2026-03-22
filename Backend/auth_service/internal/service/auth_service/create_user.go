package auth_service

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/models"
	repoModels "github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/models"
	"github.com/Artmoond/Minion-Team-TTK-Case/lib/validator"
	"golang.org/x/crypto/bcrypt"
)

const defaultRoleName = "Пользователь"

func (s *authService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {

	err := validator.Validate.Struct(req)
	if err != nil {
		log.Println("Invalid arguments: ", err.Error())
		return nil, custom_err.ErrInvalidArguments
	}

	hashPasswords, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, custom_err.ErrCreateUser
	}

	assignedRoles := []string{defaultRoleName}

	res, err := s.repo.CreateUser(ctx, &repoModels.CreateUserRequest{
		Login:      req.Login,
		Password:   hashPasswords,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		Roles:      assignedRoles,
	})

	if err != nil {
		return nil, err
	}

	token, err := s.token.GenerateToken(res.ID, assignedRoles)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{
		Token: token,
	}, err
}
