package vo

import (
	"js_statistics/app/models"
	"time"
)

type FakerReq struct {
	// 文本内容
	Text string `json:"text"`
	// 文件上传成功返回的id
	ObjID string `json:"obj_id"`
	// 0: 文本；1:图片; 2: mp3; 3:mp4
	Type int `json:"type"`
	// 文本类型下: 请求类型设置，0:text/html,1:text/plain;2:text/xml,3:application/json
	ReqType int `json:"req_type"`
	// js id
	JsID int64 `json:"js_id"`
	// 状态
	Status bool `json:"status"`
}

type FakerResp struct {
	// 文本内容
	Text string `json:"text"`
	// 文件上传成功返回的id
	ObjID string `json:"obj_id"`
	// id
	ID int64 `json:"id"`
	// 0: 文本；1:图片; 2: mp3; 3:mp4
	Type int `json:"type"`
	// 文本类型下: 请求类型设置，0:text/html,1:text/plain;2:text/xml,3:application/json
	ReqType int `json:"req_type"`
	// js id
	JsID int64 `json:"js_id"`
	// 状态
	Status bool `json:"status"`
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
	// 文本内容
	Text string `json:"text"`
	// 文件上传成功返回的id
	ObjID string `json:"obj_id"`
	// 0: 文本；1:图片; 2: mp3; 3:mp4
	Type int `json:"type"`
	// 文本类型下: 请求类型设置，0:text/html,1:text/plain;2:text/xml,3:application/json
	ReqType int `json:"req_type"`
	// js id
	JsID int64 `json:"js_id"`
	// 状态
	Status bool `json:"status"`
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
