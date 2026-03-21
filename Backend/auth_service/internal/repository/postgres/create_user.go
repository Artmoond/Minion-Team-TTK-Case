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
	roleNames := uniqueStrings(req.Roles)

	if len(roleNames) == 0 {
		log.Print("Error while creating user: no roles provided")
		return nil, custom_err.ErrCreateUser
	}

	tx, err := p.db.Begin(ctx)
	if err != nil {
		log.Printf("Error while starting transaction: %v", err)
		return nil, custom_err.ErrCreateUser
	}
	defer tx.Rollback(ctx)

	builder := sqr.Insert(tableName).Columns(loginColumn, passwordColumn, firstNameColumn, lastNameColumn, middleNameColumn).
		PlaceholderFormat(sqr.Dollar).Values(req.Login, req.Password, req.FirstName, req.LastName, req.MiddleName).Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("Error while building query: %v", err)
		return nil, custom_err.ErrBuildingQuery
	}

	err = tx.QueryRow(ctx, query, args...).Scan(&resp.ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, custom_err.ErrUserIsExist
		}

		log.Printf("Error while creating user: %v", err)
		return nil, custom_err.ErrCreateUser
	}

	roleQuery, roleArgs, err := sqr.Select(idColumn, roleNameColumn).
		From(rolesTableName).
		Where(sqr.Eq{roleNameColumn: roleNames}).
		PlaceholderFormat(sqr.Dollar).
		ToSql()
	if err != nil {
		log.Printf("Error while building role query: %v", err)
		return nil, custom_err.ErrBuildingQuery
	}

	rows, err := tx.Query(ctx, roleQuery, roleArgs...)
	if err != nil {
		log.Printf("Error while finding roles: %v", err)
		return nil, custom_err.ErrCreateUser
	}
	defer rows.Close()

	roleIDsByName := make(map[string]int64, len(roleNames))
	for rows.Next() {
		var (
			roleID   int64
			roleName string
		)

		if err = rows.Scan(&roleID, &roleName); err != nil {
			log.Printf("Error while scanning roles: %v", err)
			return nil, custom_err.ErrCreateUser
		}

		roleIDsByName[roleName] = roleID
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error while iterating roles: %v", err)
		return nil, custom_err.ErrCreateUser
	}

	if len(roleIDsByName) != len(roleNames) {
		log.Printf("Error while finding roles: requested=%v found=%v", roleNames, len(roleIDsByName))
		return nil, custom_err.ErrCreateUser
	}

	linkBuilder := sqr.Insert(userRolesTableName).
		Columns(userIDColumn, roleIDColumn).
		PlaceholderFormat(sqr.Dollar)

	for _, roleName := range roleNames {
		linkBuilder = linkBuilder.Values(resp.ID, roleIDsByName[roleName])
	}

	linkQuery, linkArgs, err := linkBuilder.ToSql()
	if err != nil {
		log.Printf("Error while building user role query: %v", err)
		return nil, custom_err.ErrBuildingQuery
	}

	if _, err = tx.Exec(ctx, linkQuery, linkArgs...); err != nil {
		log.Printf("Error while linking user role: %v", err)
		return nil, custom_err.ErrCreateUser
	}

	if err = tx.Commit(ctx); err != nil {
		log.Printf("Error while committing transaction: %v", err)
		return nil, custom_err.ErrCreateUser
	}

	resp.Roles = append(resp.Roles, roleNames...)

	return &resp, nil
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	unique := make([]string, 0, len(values))

	for _, value := range values {
		if _, exists := seen[value]; exists {
			continue
		}

		seen[value] = struct{}{}
		unique = append(unique, value)
	}

	return unique
}
