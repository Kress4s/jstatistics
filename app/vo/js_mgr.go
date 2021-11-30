package vo

import (
	"js_statistics/app/models"
	"js_statistics/types"
	"math/rand"
	"time"
)

type JsManageReq struct {
	KeyWord       string            `json:"key_word"`
	ShieldArea    string            `json:"shield_area"`
	HrefID        string            `json:"href_id"`
	Title         string            `json:"title"`
	ClientType    types.BigintArray `json:"client_type"`
	SearchEngines types.BigintArray `json:"search_engines"`
	FromMode      int               `json:"from_mode"`
	ReleaseTime   int               `json:"release_time"`
	RedirectMode  int               `json:"redirect_mode"`
	RedirectCode  int               `json:"redirect_code"`
	RedirectCount int               `json:"redirect_count"`
	WaitTime      int               `json:"wait_time"`
	CategoryID    int64             `json:"category_id"`
	Status        bool              `json:"status"`
}

func (jm *JsManageReq) ToModel(openID string) *models.JsManage {
	now := time.Now()
	sign := GenerateJSite(17)
	return &models.JsManage{
		Title:         jm.Title,
		ShieldArea:    jm.ShieldArea,
		ClientType:    jm.ClientType,
		RedirectCount: jm.RedirectCount,
		ReleaseTime:   jm.ReleaseTime,
		FromMode:      jm.FromMode,
		KeyWord:       jm.KeyWord,
		SearchEngines: jm.SearchEngines,
		RedirectMode:  jm.RedirectMode,
		RedirectCode:  jm.RedirectCode,
		HrefID:        jm.HrefID,
		WaitTime:      jm.WaitTime,
		CategoryID:    jm.CategoryID,
		Status:        jm.Status,
		Sign:          sign,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type JsManageResp struct {
	KeyWord       string            `json:"key_word"`
	Title         string            `json:"title"`
	HrefID        string            `json:"href_id"`
	ShieldArea    string            `json:"shield_area"`
	ClientType    types.BigintArray `json:"client_type"`
	SearchEngines types.BigintArray `json:"search_engines"`
	ID            int64             `json:"id"`
	FromMode      int               `json:"from_mode"`
	ReleaseTime   int               `json:"release_time"`
	RedirectCount int               `json:"redirect_count"`
	RedirectMode  int               `json:"redirect_mode"`
	RedirectCode  int               `json:"redirect_code"`
	IP            int               `json:"ip"`
	WaitTime      int               `json:"wait_time"`
	CategoryID    int64             `json:"category_id"`
	Status        bool              `json:"status"`
}

func NewJsManageResponse(jm *models.JsManage) *JsManageResp {
	return &JsManageResp{
		ID:    jm.ID,
		Title: jm.Title,
		//TODO:统计
		IP:            0,
		ShieldArea:    jm.ShieldArea,
		ClientType:    jm.ClientType,
		RedirectCount: jm.RedirectCount,
		ReleaseTime:   jm.ReleaseTime,
		FromMode:      jm.FromMode,
		KeyWord:       jm.KeyWord,
		SearchEngines: jm.SearchEngines,
		RedirectMode:  jm.RedirectMode,
		RedirectCode:  jm.RedirectCode,
		HrefID:        jm.HrefID,
		WaitTime:      jm.WaitTime,
		CategoryID:    jm.CategoryID,
		Status:        jm.Status,
	}
}

type JsManageUpdateReq struct {
	ShieldArea    string            `json:"shield_area"`
	HrefID        string            `json:"href_id"`
	KeyWord       string            `json:"key_word"`
	ClientType    types.BigintArray `json:"client_type"`
	SearchEngines types.BigintArray `json:"search_engines"`
	RedirectCount int               `json:"redirect_count"`
	FromMode      int               `json:"from_mode"`
	RedirectMode  int               `json:"redirect_mode"`
	RedirectCode  int               `json:"redirect_code"`
	ReleaseTime   int               `json:"release_time"`
	WaitTime      int               `json:"wait_time"`
	CategoryID    int64             `json:"category_id"`
}

func (jum *JsManageUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"shield_area":    jum.ShieldArea,
		"client_type":    jum.ClientType,
		"redirect_count": jum.RedirectCount,
		"release_time":   jum.ReleaseTime,
		"from_mode":      jum.FromMode,
		"key_word":       jum.KeyWord,
		"search_engines": jum.SearchEngines,
		"redirect_mode":  jum.RedirectMode,
		"redirect_code":  jum.RedirectCode,
		"href_id":        jum.HrefID,
		"wait_time":      jum.WaitTime,
		"category_id":    jum.CategoryID,
		"update_by":      openID,
		"update_at":      time.Now(),
	}
}

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateJSite(n int) string {
	return randSeq(n)
}

type JSiteResp struct {
	Site string `json:"site"`
}
