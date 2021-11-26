package vo

import (
	"js_statistics/app/models"
	"time"
)

type FakerReq struct {
	// 类型 0:文本；1:图片 2：音频 3：视频
	Type int `json:"type"`
	// 请求类型(type为文本情况下) 0:text/html,1:text/plain;2:text/xml,3:application/json
	ReqType int `json:"req_type"`
	// 文本内容
	Text string `json:"text"`
	// 上传文件接口返回的id
	ObjID string `json:"obj_id"`
	// 开启状态
	Status bool `json:"status"`
}

type FakerResp struct {
	// id
	ID int64 `json:"id"`
	// 类型 0:文本；1:图片 2：音频 3：视频
	Type int `json:"type"`
	// 请求类型(type为文本情况下) 0:text/html,1:text/plain;2:text/xml,3:application/json
	ReqType int `json:"req_type"`
	// 文本内容
	Text string `json:"text"`
	// 上传文件接口返回的id
	ObjID string `json:"obj_id"`
	// 开启状态
	Status bool `json:"status"`
}

func NewFakerResponse(f *models.Faker) *FakerResp {
	return &FakerResp{
		ID:      f.ID,
		Type:    f.Type,
		ReqType: f.ReqType,
		Text:    f.Text,
		ObjID:   f.ObjID,
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
	// 类型 0:文本；1:图片 2：音频 3：视频
	Type int `json:"type"`
	// 请求类型(type为文本情况下) 0:text/html,1:text/plain;2:text/xml,3:application/json
	ReqType int `json:"req_type"`
	// 文本内容
	Text string `json:"text"`
	// 上传文件接口返回的id
	ObjID string `json:"obj_id"`
	// 开启状态
	Status bool `json:"status"`
}

func (fuq *FakerUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"type":     fuq.Type,
		"req_type": fuq.ReqType,
		"text":     fuq.Text,
		"obj_id":   fuq.ObjID,
		"status":   fuq.Status,
	}
}
