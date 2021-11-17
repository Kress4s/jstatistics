package vo

import (
	"js_statistics/app/models"
	"time"
)

type CDNReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (cr *CDNReq) ToModel(openID string) *models.CDN {
	now := time.Now()
	return &models.CDN{
		Title: cr.Title,
		IP:    cr.IP,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type CDNResp struct {
	// ID
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// cdn
	IP string `json:"ip"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
}

func NewCDNResponse(im *models.CDN) *CDNResp {
	return &CDNResp{
		ID:       im.ID,
		Title:    im.Title,
		IP:       im.IP,
		CreateAt: im.CreateAt,
	}
}

type CDNUpdateReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (cur *CDNUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":     cur.Title,
		"ip":        cur.IP,
		"update_by": openID,
		"update_at": time.Now(),
	}
}
