package vo

import "js_statistics/app/models"

type UserUpdateRolesReq struct {
	// 角色IDS
	RoleIDs []int64 `json:"role_ids,omitempty"`
}

func (urq *UserUpdateRolesReq) ToModel(openID string, uid int64) []models.UserRoleRelation {
	rolesID := urq.RoleIDs
	ur := make([]models.UserRoleRelation, 0, len(rolesID))
	for i := range rolesID {
		ur = append(ur, models.UserRoleRelation{
			UserID: uid,
			RoleID: rolesID[i],
			Base: models.Base{
				CreateBy: openID,
				UpdateBy: openID,
			},
		})
	}
	return ur
}
