package postgres

import (
	"context"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository/models"
)

func (p *Postgres) UpdateUserRole(ctx context.Context, req *models.UpdateUserRoleReq) error {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, r := range req.Rol {
		var roleID int64
		err := tx.QueryRow(ctx,
			`SELECT id FROM roles WHERE rolename = $1`,
			r.Role,
		).Scan(&roleID)

		if err != nil {
			return err
		}

		if r.IsAccept {
			_, err = tx.Exec(ctx,
				`INSERT INTO user_roles (user_id, role_id)
				 VALUES ($1, $2)
				 ON CONFLICT DO NOTHING`,
				req.ID, roleID,
			)
		} else {
			_, err = tx.Exec(ctx,
				`DELETE FROM user_roles
				 WHERE user_id = $1 AND role_id = $2`,
				req.ID, roleID,
			)
		}

		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
