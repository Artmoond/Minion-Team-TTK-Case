package postgres

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/models"
	sqr "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (p *postgres) FindByName(ctx context.Context, login string) (*models.FindByNameResponse, error) {
	var resp models.FindByNameResponse

	builder := sqr.
		Select(
			"u."+idColumn,
			"u."+loginColumn,
			"u."+passwordColumn,
			"COALESCE(array_agg(r."+roleNameColumn+") FILTER (WHERE r."+roleNameColumn+" IS NOT NULL), '{}') AS roles",
		).
		From(tableName + " u").
		LeftJoin(userRolesTableName + " ur ON ur." + userIDColumn + " = u." + idColumn).
		LeftJoin(rolesTableName + " r ON r." + idColumn + " = ur." + roleIDColumn).
		Where(sqr.Eq{"u." + loginColumn: login}).
		GroupBy("u." + idColumn, "u." + loginColumn, "u." + passwordColumn).
		PlaceholderFormat(sqr.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("Error while building query: %v", err)
		return nil, custom_err.ErrBuildingQuery
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&resp.ID, &resp.Login, &resp.Password, &resp.Roles)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, custom_err.ErrUserNotFound
		}

		log.Printf("Error while finding user by login: %v", err)
		return nil, custom_err.ErrCreateUser
	}

	return &resp, nil
}
