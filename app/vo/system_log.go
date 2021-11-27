package vo

import (
	"js_statistics/app/models"
	"time"
)

type SystemLogReq struct {
	UserName    string `json:"user_name"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func (slr *SystemLogReq) ToModel() *models.SystemLog {
	return &models.SystemLog{
		UserName:    slr.UserName,
		IP:          slr.IP,
		Address:     slr.Address,
		Description: slr.Description,
		OperateAt:   time.Now(),
	}
}

type SystemLogResp struct {
	ID          int64     `json:"id"`
	UserName    string    `json:"user_name"`
	IP          string    `json:"ip"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	OperateAt   time.Time `json:"operateAt"`
}

func NewSysLogResp(log *models.SystemLog) *SystemLogResp {
	return &SystemLogResp{
		ID:          log.ID,
		UserName:    log.UserName,
		IP:          log.IP,
		Address:     log.Address,
		Description: log.Description,
		OperateAt:   log.OperateAt,
	}
}
