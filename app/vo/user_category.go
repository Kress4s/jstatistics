package vo

import "js_statistics/app/models"

type UserUpdateJscAndJsReq struct {
	// js分类IDS
	CategoryIDs []int64 `json:"category_ids"`
	// js主分类ID
	PrimaryIDs []int64 `json:"primary_ids"`
}

func (urq *UserUpdateJscAndJsReq) ToModel(openID string, uid int64) ([]models.UserCategoryRelation,
	[]models.UserPrimaryRelation) {
	ur := make([]models.UserCategoryRelation, 0, len(urq.CategoryIDs))
	up := make([]models.UserPrimaryRelation, 0, len(urq.PrimaryIDs))
	for i := range urq.CategoryIDs {
		ur = append(ur, models.UserCategoryRelation{
			UserID:     uid,
			CategoryID: urq.CategoryIDs[i],
			Base: models.Base{
				CreateBy: openID,
				UpdateBy: openID,
			},
		})
	}
	for i := range urq.PrimaryIDs {
		up = append(up, models.UserPrimaryRelation{
			UserID:    uid,
			PrimaryID: urq.PrimaryIDs[i],
			Base: models.Base{
				CreateBy: openID,
				UpdateBy: openID,
			},
		})
	}
	return ur, up
}
