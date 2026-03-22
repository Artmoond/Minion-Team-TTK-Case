package postgres

import (
	"context"
	"log"
	"time"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	sqr "github.com/Masterminds/squirrel"
)

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	builder := sqr.Update(tableName).Set(deletedColumn, time.Now().UTC()).Where(sqr.Eq{idColumn: id}).PlaceholderFormat(sqr.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Error create sql")
		return custom_err.ErrBuildingQuery
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		log.Println("Error deleting user", id, err)
		return custom_err.ErrBuildingQuery
	}

	return nil
}
