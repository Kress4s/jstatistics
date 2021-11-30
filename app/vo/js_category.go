package vo

import (
	"js_statistics/app/models"
	"time"
)

type JsCategoryReq struct {
	Title     string `json:"title"`
	PrimaryID int64  `json:"primary_id"`
	DomainID  int64  `json:"domain_id"`
}

func (jpr *JsCategoryReq) ToModel(openID string) *models.JsCategory {
	now := time.Now()
	return &models.JsCategory{
		DomainID:  jpr.DomainID,
		Title:     jpr.Title,
		PrimaryID: jpr.PrimaryID,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type Domain struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

type JsCategoryResp struct {
	Domain    *Domain       `json:"domain"`
	Title     string        `json:"title"`
	JsPrimary JsPrimaryResp `json:"primary"`
	ID        int64         `json:"id"`
}

func NewJsCategoryResponse(jp *models.JsCategory, domain *models.DomainMgr, jsp *models.JsPrimary) *JsCategoryResp {
	domainModel := &Domain{}
	if domain != nil {
		domainModel.ID = domain.ID
		domainModel.Title = domain.Title
	} else {
		domainModel = nil
	}
	return &JsCategoryResp{
		ID:        jp.ID,
		Title:     jp.Title,
		Domain:    domainModel,
		JsPrimary: *NewJsPrimaryResponse(jsp),
	}
}

type JsCategoryUpdateReq struct {
	Title     string `json:"title"`
	PrimaryID int64  `json:"primary_id"`
	DomainID  int64  `json:"domain_id"`
}

func (jpr *JsCategoryUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":      jpr.Title,
		"primary_id": jpr.PrimaryID,
		"domain_id":  jpr.DomainID,
		"update_by":  openID,
		"update_at":  time.Now(),
	}
}
