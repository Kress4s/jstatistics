package vo

import (
	"js_statistics/app/models"
	"time"
)

type FakerReq struct {
	Text    string `json:"text"`
	ObjID   string `json:"obj_id"`
	Type    int    `json:"type"`
	ReqType int    `json:"req_type"`
	JsID    int64  `json:"js_id"`
	Status  bool   `json:"status"`
}

type FakerResp struct {
	Text    string `json:"text"`
	ObjID   string `json:"obj_id"`
	ID      int64  `json:"id"`
	Type    int    `json:"type"`
	ReqType int    `json:"req_type"`
	JsID    int64  `json:"js_id"`
	Status  bool   `json:"status"`
}

func NewFakerResponse(f *models.Faker) *FakerResp {
	return &FakerResp{
		ID:      f.ID,
		Type:    f.Type,
		ReqType: f.ReqType,
		Text:    f.Text,
		ObjID:   f.ObjID,
		JsID:    f.JsID,
		Status:  f.Status,
	}
}

func (fr *FakerReq) ToModel(openID string) *models.Faker {
	now := time.Now()
	return &models.Faker{
		Type:    fr.Type,
		ReqType: fr.ReqType,
		Text:    fr.Text,
		ObjID:   fr.ObjID,
		JsID:    fr.JsID,
		Status:  fr.Status,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type FakerUpdateReq struct {
	Text    string `json:"text"`
	ObjID   string `json:"obj_id"`
	Type    int    `json:"type"`
	ReqType int    `json:"req_type"`
	JsID    int64  `json:"js_id"`
	Status  bool   `json:"status"`
}

func (fuq *FakerUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"type":     fuq.Type,
		"req_type": fuq.ReqType,
		"text":     fuq.Text,
		"obj_id":   fuq.ObjID,
		"js_id":    fuq.JsID,
		"status":   fuq.Status,
	}
}
