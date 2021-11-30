package vo

import (
	"js_statistics/app/models"
	"time"
)

type RedirectManageReq struct {
	OFF        time.Time `json:"off,omitempty"`
	ON         time.Time `json:"on,omitempty"`
	PC         string    `json:"pc"`
	Android    string    `json:"android"`
	IOS        string    `json:"ios"`
	Title      string    `json:"title"`
	CategoryID int64     `json:"category_id"`
}

func (rmr *RedirectManageReq) ToModel(openID string) *models.RedirectManage {
	return &models.RedirectManage{
		Title:      rmr.Title,
		PC:         rmr.PC,
		Android:    rmr.Android,
		IOS:        rmr.IOS,
		CategoryID: rmr.CategoryID,
		ON:         rmr.ON,
		OFF:        rmr.OFF,
	}
}

type RedirectManageResp struct {
	OFF     time.Time `json:"off"`
	ON      time.Time `json:"on"`
	Title   string    `json:"title"`
	PC      string    `json:"pc"`
	Android string    `json:"android"`
	IOS     string    `json:"ios"`
	ID      int64     `json:"id"`
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
	OFF        time.Time `json:"off"`
	ON         time.Time `json:"on"`
	PC         string    `json:"pc"`
	Android    string    `json:"android"`
	IOS        string    `json:"ios"`
	Title      string    `json:"title"`
	CategoryID int64     `json:"category_id"`
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
	UpdateAt   time.Time `json:"update_at"`
	OldPC      string    `json:"old_pc"`
	OldAndroid string    `json:"old_android"`
	OldIOS     string    `json:"old_ios"`
	PC         string    `json:"pc"`
	IOS        string    `json:"ios"`
	Type       string    `json:"type"`
	Android    string    `json:"android"`
	ID         int64     `json:"id"`
	RedirectID int64     `json:"redirect_id"`
	CategoryID int64     `json:"category_id"`
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
