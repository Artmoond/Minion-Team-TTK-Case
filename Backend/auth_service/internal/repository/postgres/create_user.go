package postgres

import (
	"context"
	"log"
	"strings"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/models"
	sqr "github.com/Masterminds/squirrel"
)

func (p *postgres) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	var resp models.CreateUserResponse

	builder := sqr.Insert(tableName).Columns(loginColumn, passwordColumn, firstNameColumn, lastNameColumn, middleNameColumn, roleColumn).
		PlaceholderFormat(sqr.Dollar).Values(req.Login, req.Password, req.FirstName, req.LastName, req.MiddleName, req.Role).Suffix("RETURNING id, role")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("Error while building query: %v", err)
		return nil, custom_err.ErrBuildingQuery
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&resp.ID, &resp.Role)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, custom_err.ErrUserIsExist
		}

		log.Printf("Error while creating user: %v", err)
		return nil, custom_err.ErrCreateUser
	}

	return &resp, nil
}
