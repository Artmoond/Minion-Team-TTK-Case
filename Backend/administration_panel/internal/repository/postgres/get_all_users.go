package postgres

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/converter"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	serviceModels "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository/models"
	sqr "github.com/Masterminds/squirrel"
)

func (p *Postgres) GetAllUsers(ctx context.Context) ([]serviceModels.User, error) {
	res := make([]models.User, 0)

	builder := sqr.Select(
		"u.id",
		"u.login",
		"u.first_name",
		"u.last_name",
		"u.middle_name",
		"u.date",
		"COALESCE(array_agg(r.rolename) FILTER (WHERE r.rolename IS NOT NULL), '{}') AS roles",
	).From("users u").
		LeftJoin("user_roles ur ON ur.user_id = u.id").
		LeftJoin("roles r ON r.id = ur.role_id").
		GroupBy(
			"u.id",
			"u.login",
			"u.first_name",
			"u.last_name",
			"u.middle_name",
			"u.date",
		).
		OrderBy("u.id ASC").
		PlaceholderFormat(sqr.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Error while building role query")
		return nil, custom_err.ErrBuildingQuery
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("Error while rows query")
		return nil, custom_err.ErrGetAllUsers
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Date,
			&user.Roles,
		)
		if err != nil {
			log.Println("Error while scanning rows: ", err)
			return nil, custom_err.ErrGetAllUsers
		}

		res = append(res, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return converter.ToServiceFromREPO(res), nil
}
