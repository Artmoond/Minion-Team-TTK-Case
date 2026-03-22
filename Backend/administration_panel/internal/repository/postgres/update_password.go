package postgres

import (
	"context"
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	sqr "github.com/Masterminds/squirrel"
)

func (p *Postgres) UpdatePassword(ctx context.Context, id int64, newPassword []byte) error {
	builder := sqr.Update(tableName).Where(sqr.Eq{idColumn: id}).Set(passwordColumn, newPassword).PlaceholderFormat(sqr.Dollar)
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
