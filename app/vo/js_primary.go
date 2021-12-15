package vo

import (
	"js_statistics/app/models"
	"time"
)

type JsPrimaryReq struct {
	// 标题
	Title string `json:"title"`
}

func (jpr *JsPrimaryReq) ToModel(openID string) *models.JsPrimary {
	now := time.Now()
	return &models.JsPrimary{
		Title: jpr.Title,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type JsPrimaryResp struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

func NewJsPrimaryResponse(jp *models.JsPrimary) *JsPrimaryResp {
	return &JsPrimaryResp{
		ID:    jp.ID,
		Title: jp.Title,
	}
}

type JsPrimaryUpdateReq struct {
	// 标题
	Title string `json:"title"`
}

func (jpr *JsPrimaryUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":     jpr.Title,
		"update_by": openID,
		"update_at": time.Now(),
	}
}

type Primaries struct {
	Title      string            `json:"title"`
	Categories []JsCategoryBrief `json:"categories"`
	ID         int64             `json:"id"`
}

type PrimaryKey struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}
