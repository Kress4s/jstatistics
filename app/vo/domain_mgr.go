package vo

import (
	"js_statistics/app/models"
	"time"
)

type DomainReq struct {
	// 标题
	Title string `json:"title"`
	// 域名
	Domain string `json:"domain"`
	// ssl
	SSL bool `json:"ssl,omitempty"`
	// 证书
	Certificate string `json:"certificate,omitempty"`
	// 秘钥
	SecretKey string `json:"secret_key,omitempty"`
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
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Domain      string `json:"domain"`
	SSL         bool   `json:"ssl"`
	Certificate string `json:"certificate"`
	SecretKey   string `json:"secret_key"`
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
	// 标题
	Title string `json:"title"`
	// 域名
	Domain string `json:"domain"`
	// ssl
	SSL bool `json:"ssl,omitempty"`
	// 证书
	Certificate string `json:"certificate,omitempty"`
	// 秘钥
	SecretKey string `json:"secret_key,omitempty"`
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
