package vo

import (
	"js_statistics/app/models"
	"time"
)

type DomainReq struct {
	Title       string `json:"title"`
	Domain      string `json:"domain"`
	Certificate string `json:"certificate,omitempty"`
	SecretKey   string `json:"secret_key,omitempty"`
	SSL         bool   `json:"ssl,omitempty"`
}

func (dr *DomainReq) ToModel(openID string) *models.DomainMgr {
	now := time.Now()
	return &models.DomainMgr{
		Title:       dr.Title,
		Domain:      dr.Domain,
		SSL:         dr.SSL,
		Certificate: dr.Certificate,
		SecretKey:   dr.SecretKey,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type DomainResp struct {
	Title       string `json:"title"`
	Domain      string `json:"domain"`
	Certificate string `json:"certificate"`
	SecretKey   string `json:"secret_key"`
	ID          int64  `json:"id"`
	SSL         bool   `json:"ssl"`
}

func NewDomainResponse(dm *models.DomainMgr) *DomainResp {
	return &DomainResp{
		ID:          dm.ID,
		Title:       dm.Title,
		Domain:      dm.Domain,
		SSL:         dm.SSL,
		Certificate: dm.Certificate,
		SecretKey:   dm.SecretKey,
	}
}

type DomainUpdateReq struct {
	Title       string `json:"title"`
	Domain      string `json:"domain"`
	Certificate string `json:"certificate,omitempty"`
	SecretKey   string `json:"secret_key,omitempty"`
	SSL         bool   `json:"ssl,omitempty"`
}

func (dur *DomainUpdateReq) ToMap(openID string) map[string]interface{} {
	now := time.Now()
	return map[string]interface{}{
		"title":       dur.Title,
		"domain":      dur.Domain,
		"ssl":         dur.SSL,
		"certificate": dur.Certificate,
		"secret_key":  dur.SecretKey,
		"update_by":   openID,
		"update_at":   now,
	}
}
