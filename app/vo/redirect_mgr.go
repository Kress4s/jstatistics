package vo

import (
	"js_statistics/app/models"
	"time"
)

type RedirectManageReq struct {
	// 标题
	Title string `json:"title"`
	// PC端跳转地址
	PC string `json:"pc"`
	// android端跳转地址
	Android string `json:"android"`
	// ios跳转地址
	IOS string `json:"ios"`
	// 开启时间
	ON time.Time `json:"on"`
	// 关闭时间
	OFF time.Time `json:"off"`
}

func (rmr *RedirectManageReq) ToModel(openID string) *models.RedirectManage {
	return &models.RedirectManage{
		Title:   rmr.Title,
		PC:      rmr.PC,
		Android: rmr.Android,
		IOS:     rmr.IOS,
		ON:      rmr.ON,
		OFF:     rmr.OFF,
	}
}

type RedirectManageResp struct {
	// id
	ID uint `json:"id"`
	// 标题
	Title string `json:"title"`
	// PC端跳转地址
	PC string `json:"pc"`
	// android端跳转地址
	Android string `json:"android"`
	// ios跳转地址
	IOS string `json:"ios"`
	// 开启时间
	ON time.Time `json:"on"`
	// 关闭时间
	OFF time.Time `json:"off"`
}

func NewRedirectManageResponse(rm *models.RedirectManage) *RedirectManageResp {
	return &RedirectManageResp{
		ID:      rm.ID,
		Title:   rm.Title,
		PC:      rm.PC,
		Android: rm.Android,
		IOS:     rm.IOS,
		ON:      rm.ON,
		OFF:     rm.OFF,
	}
}

type RedirectManageUpdateReq struct {
	// 标题
	Title string `json:"title"`
	// PC端跳转地址
	PC string `json:"pc"`
	// android端跳转地址
	Android string `json:"android"`
	// ios跳转地址
	IOS string `json:"ios"`
	// 开启时间
	ON time.Time `json:"on"`
	// 关闭时间
	OFF time.Time `json:"off"`
}

func (rmr *RedirectManageUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":   rmr.Title,
		"pc":      rmr.PC,
		"android": rmr.Android,
		"ios":     rmr.IOS,
		"on":      rmr.ON,
		"off":     rmr.OFF,
	}
}
