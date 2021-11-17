package vo

import (
	"js_statistics/app/models"
	"time"
)

type JsCategoryReq struct {
	// js主分类ID
	PrimaryID int64 `json:"primary_id"`
	// 标题
	Title string `json:"title"`
	// 域名ID
	DomainID int64 `json:"domain_id"`
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
	// 域名配置的id
	ID int64 `json:"id"`
	// 域名配置的标题
	Title string `json:"title"`
}

type JsCategoryResp struct {
	// id
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 域名配置信息
	Domain *Domain `json:"domain"`
	// js主分类的信息
	JsPrimary JsPrimaryResp `json:"primary"`
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
	// js主分类ID
	PrimaryID int64 `json:"primary_id"`
	// 标题
	Title string `json:"title"`
	// 域名ID
	DomainID int64 `json:"domain_id"`
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
