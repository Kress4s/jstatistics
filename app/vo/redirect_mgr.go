package vo

import (
	"js_statistics/app/models"
	"time"
)

type RedirectManageReq struct {
	// 关闭时间(不选不传); eg: 01:02:03
	OFF *string `json:"off,omitempty"`
	// 开启时间(不选不传); eg: 01:02:03
	ON *string `json:"on,omitempty"`
	// pc跳转地址
	PC string `json:"pc"`
	// android跳转地址
	Android string `json:"android"`
	// ios跳转地址
	IOS string `json:"ios"`
	// 标题
	Title string `json:"title"`
	// 分类id
	CategoryID int64 `json:"category_id"`
}

func (rmr *RedirectManageReq) ToModel(openID string) *models.RedirectManage {
	return &models.RedirectManage{
		Title:      rmr.Title,
		PC:         rmr.PC,
		Android:    rmr.Android,
		IOS:        rmr.IOS,
		CategoryID: rmr.CategoryID,
		Status:     true,
		ON:         rmr.ON,
		OFF:        rmr.OFF,
	}
}

type RedirectManageResp struct {
	// 开启时间(不选不传)
	OFF *string `json:"off"`
	// 开启时间(不选不传)
	ON      *string `json:"on"`
	Title   string  `json:"title"`
	PC      string  `json:"pc"`
	Android string  `json:"android"`
	IOS     string  `json:"ios"`
	ID      int64   `json:"id"`
	Status  bool    `json:"status"`
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
		Status:  rm.Status,
	}
}

type RedirectManageUpdateReq struct {
	OFF     *string `json:"off"`
	ON      *string `json:"on"`
	PC      string  `json:"pc"`
	Android string  `json:"android"`
	IOS     string  `json:"ios"`
	Title   string  `json:"title"`
	// js分类id
	CategoryID int64 `json:"category_id"`
}

func (rmr *RedirectManageUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":       rmr.Title,
		"pc":          rmr.PC,
		"android":     rmr.Android,
		"ios":         rmr.IOS,
		"category_id": rmr.CategoryID,
		"on":          rmr.ON,
		"off":         rmr.OFF,
	}
}

func RedirectLog(rm *models.RedirectManage) *models.RedirectLog {
	return &models.RedirectLog{
		RedirectID: rm.ID,
		CategoryID: rm.CategoryID,
		PC:         rm.PC,
		Android:    rm.Android,
		IOS:        rm.IOS,
		UpdateAt:   time.Now(),
		Type:       "新建",
	}
}

type RedirectLogResp struct {
	// 修改时间
	UpdateAt time.Time `json:"update_at"`
	// 原pc跳转地址
	OldPC string `json:"old_pc"`
	// 原android跳转地址
	OldAndroid string `json:"old_android"`
	// 原IOS跳转地址
	OldIOS string `json:"old_ios"`
	// 现pc跳转地址
	PC string `json:"pc"`
	// 现ios跳转地址
	IOS string `json:"ios"`
	// 操作类型
	Type string `json:"type"`
	// 现android跳转地址
	Android string `json:"android"`
	// ID
	ID int64 `json:"id"`
	// 跳转管理id
	RedirectID int64 `json:"redirect_id"`
	// js分类id
	CategoryID int64 `json:"category_id"`
}

func NewRedirectLogResp(rl *models.RedirectLog) *RedirectLogResp {
	return &RedirectLogResp{
		ID:         rl.ID,
		RedirectID: rl.RedirectID,
		CategoryID: rl.CategoryID,
		PC:         rl.PC,
		Android:    rl.Android,
		IOS:        rl.IOS,
		OldPC:      rl.OldPC,
		OldAndroid: rl.OldAndroid,
		OldIOS:     rl.OldIOS,
		Type:       rl.Type,
		UpdateAt:   rl.UpdateAt,
	}
}
