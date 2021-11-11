package vo

import (
	"js_statistics/app/models"
)

type RolePermissionReq struct {
	RoleID uint
}

func (rpq *RolePermissionReq) ToModel(openID string, pmsids []uint) ([]models.RolePermissionRelation, error) {
	ps := make([]models.RolePermissionRelation, 0, len(pmsids))
	for i := range pmsids {
		ps = append(ps, models.RolePermissionRelation{
			RoleID:       rpq.RoleID,
			PermissionID: pmsids[i],
			Base: models.Base{
				CreateBy: openID,
				UpdateBy: openID,
			},
		})
	}
	return ps, nil
}
