package vo

import (
	"js_statistics/app/models"
	"js_statistics/types"
	"time"
)

type JsManageReq struct {
	// 标题
	Title string `json:"title"`
	// 屏蔽地区，多个用 ”-“ 相连；eg：北京市-上海市-...
	ShieldArea string `json:"shield_area"`
	// 客户端 0：移动端； 1：PC端
	ClientType types.BigintArray `json:"client_type"`
	// 跳转次数
	RedirectCount int `json:"redirect_count"`
	// 封禁小时
	ReleaseTime int `json:"release_time"`
	// 来源 0：无；1：关键词 2：搜索引擎
	FromMode int `json:"from_mode"`
	// 关键词
	KeyWord string `json:"key_word"`
	// 搜索引擎
	SearchEngines types.BigintArray `json:"search_engines"`
	// 跳转方式
	RedirectMode int `json:"redirect_mode"`
	// 跳转代码 0：Top；1：Windows
	RedirectCode int `json:"redirect_code"`
	// href 跳转id，多个用 "," 连接
	HrefID string `json:"href_id"`
	// 等待时间
	WaitTime int `json:"wait_time"`
	// 所属js分类的ID
	CategoryID uint `json:"category_id"`
}

func (jm *JsManageReq) ToModel(openID string) *models.JsManage {
	now := time.Now()
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
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type JsManageResp struct {
	// id
	ID uint `json:"id"`
	// 标题
	Title string `json:"title"`
	// 今日IP数
	IP int `json:"ip"`
	// 屏蔽地区，多个用 ”-“ 相连；eg：北京市-上海市-...
	ShieldArea string `json:"shield_area"`
	// 客户端 0：移动端； 1：PC端
	ClientType types.BigintArray `json:"client_type"`
	// 跳转次数
	RedirectCount int `json:"redirect_count"`
	// 封禁小时
	ReleaseTime int `json:"release_time"`
	// 来源 0：无；1：关键词 2：搜索引擎
	FromMode int `json:"from_mode"`
	// 关键词
	KeyWord string `json:"key_word"`
	// 搜索引擎
	SearchEngines types.BigintArray `json:"search_engines"`
	// 跳转方式
	RedirectMode int `json:"redirect_mode"`
	// 跳转代码 0：Top；1：Windows
	RedirectCode int `json:"redirect_code"`
	// href 跳转id，多个用 "," 连接
	HrefID string `json:"href_id"`
	// 等待时间
	WaitTime int `json:"wait_time"`
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
	}
}

type JsManageUpdateReq struct {
	// 屏蔽地区，多个用 ”-“ 相连；eg：北京市-上海市-...
	ShieldArea string `json:"shield_area"`
	// 客户端 0：移动端； 1：PC端
	ClientType types.BigintArray `json:"client_type"`
	// 跳转次数
	RedirectCount int `json:"redirect_count"`
	// 封禁小时
	ReleaseTime int `json:"release_time"`
	// 来源 0：无；1：关键词 2：搜索引擎
	FromMode int `json:"from_mode"`
	// 关键词
	KeyWord string `json:"key_word"`
	// 搜索引擎
	SearchEngines types.BigintArray `json:"search_engines"`
	// 跳转方式
	RedirectMode int `json:"redirect_mode"`
	// 跳转代码 0：Top；1：Windows
	RedirectCode int `json:"redirect_code"`
	// href 跳转id，多个用 "," 连接
	HrefID string `json:"href_id"`
	// 等待时间
	WaitTime int `json:"wait_time"`
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
		"update_by":      openID,
		"update_at":      time.Now(),
	}
}
