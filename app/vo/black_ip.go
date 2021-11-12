package vo

import (
	"js_statistics/app/models"
	"time"
)

type BlackIPReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (ir *BlackIPReq) ToModel(openID string) *models.BlackIPMgr {
	now := time.Now()
	return &models.BlackIPMgr{
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

type BlackIPResp struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title"`
	IP       string    `json:"ip"`
	CreateAt time.Time `json:"create_at"`
}

func NewBlackIPResponse(im *models.BlackIPMgr) *BlackIPResp {
	return &BlackIPResp{
		ID:       im.ID,
		Title:    im.Title,
		IP:       im.IP,
		CreateAt: im.CreateAt,
	}
}

type BlackIPUpdateReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (iur *BlackIPUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":     iur.Title,
		"ip":        iur.IP,
		"update_by": openID,
		"update_at": time.Now(),
	}
}
