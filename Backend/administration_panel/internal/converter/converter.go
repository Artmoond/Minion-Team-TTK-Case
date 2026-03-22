package converter

import (
	serviceModels "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	repoModels "github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/repository/models"
)

func ToServiceFromREPO(req []repoModels.User) []serviceModels.User {
	var res []serviceModels.User

	for _, v := range req {
		user := serviceModels.User{
			ID:         v.ID,
			Login:      v.Login,
			Roles:      v.Roles,
			FirstName:  v.FirstName,
			LastName:   v.LastName,
			MiddleName: v.MiddleName,
			Date:       v.Date,
		}

		res = append(res, user)
	}

	return res
}

func ToRepoFromService(req *serviceModels.UpdateUserRoleReq) *repoModels.UpdateUserRoleReq {
	res := repoModels.UpdateUserRoleReq{
		ID:  req.ID,
		Rol: make([]repoModels.Roles, 0, len(req.Rol)),
	}

	for _, v := range req.Rol {
		r := repoModels.Roles{
			Role:     v.Role,
			IsAccept: v.IsAccept,
		}

		res.Rol = append(res.Rol, r)
	}

	return &res
}
