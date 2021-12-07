package vo

import "js_statistics/app/models"

type UserUpdatePrimaryReq struct {
	// js分类IDS
	PrimaryIDs []int64 `json:"primary_ids"`
}

func (urq *UserUpdatePrimaryReq) ToModel(openID string, uid int64) []models.UserPrimaryRelation {
	categoriesID := urq.PrimaryIDs
	ur := make([]models.UserPrimaryRelation, 0, len(categoriesID))
	for i := range categoriesID {
		ur = append(ur, models.UserPrimaryRelation{
			UserID:    uid,
			PrimaryID: categoriesID[i],
			Base: models.Base{
				CreateBy: openID,
				UpdateBy: openID,
			},
		})
	}
	return ur
}
