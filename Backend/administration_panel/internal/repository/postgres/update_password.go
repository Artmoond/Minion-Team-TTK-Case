package repository

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/repository/postgres"
	sqr "github.com/Masterminds/squirrel"
)

func (p *postgres.postgres) UpdatePassword(ctx context.Context, id string, newPassword string) error {
	builder := sqr.Update(postgres.tableName).Where(sqr.Eq{postgres.idColumn: id}).Set(postgres.passwordColumn, newPassword).PlaceholderFormat(sqr.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("Error while building query: %v, %v", query, err)
		return custom_err.ErrBuildingQuery
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Error while executing query: %v, %v", query, err)
		return custom_err.ErrBuildingQuery
	}

	return nil
}
