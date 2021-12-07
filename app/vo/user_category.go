package vo

import "js_statistics/app/models"

type UserUpdateCategoryReq struct {
	// js分类IDS
	CategoryIDs []int64 `json:"category_ids"`
}

func (urq *UserUpdateCategoryReq) ToModel(openID string, uid int64) []models.UserCategoryRelation {
	categoriesID := urq.CategoryIDs
	ur := make([]models.UserCategoryRelation, 0, len(categoriesID))
	for i := range categoriesID {
		ur = append(ur, models.UserCategoryRelation{
			UserID:     uid,
			CategoryID: categoriesID[i],
			Base: models.Base{
				CreateBy: openID,
				UpdateBy: openID,
			},
		})
	}
	return ur
}


