package vo

import (
	"js_statistics/app/models"
	"time"
)

type IPReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (ir *IPReq) ToModel(openID string) *models.WhiteIP {
	now := time.Now()
	return &models.WhiteIP{
		Title: ir.Title,
		IP:    ir.IP,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type IPResp struct {
	ID       int64      `json:"id"`
	Title    string    `json:"title"`
	IP       string    `json:"ip"`
	CreateAt time.Time `json:"create_at"`
}

func NewIPResponse(im *models.WhiteIP) *IPResp {
	return &IPResp{
		ID:       im.ID,
		Title:    im.Title,
		IP:       im.IP,
		CreateAt: im.CreateAt,
	}
}

type IPUpdateReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (iur *IPUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":     iur.Title,
		"ip":        iur.IP,
		"update_by": openID,
		"update_at": time.Now(),
	}
}
