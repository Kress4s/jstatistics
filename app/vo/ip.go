package vo

import (
	"js_statistics/app/models"
	"time"
)

type IPReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (ir *IPReq) ToModel(openID string) *models.WhiteIP {
	now := time.Now()
	return &models.WhiteIP{
		Title: ir.Title,
		IP:    ir.IP,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}
}

type IPResp struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	IP       string    `json:"ip"`
	CreateAt time.Time `json:"create_at"`
}

func NewIPResponse(im *models.WhiteIP) *IPResp {
	return &IPResp{
		ID:       im.ID,
		Title:    im.Title,
		IP:       im.IP,
		CreateAt: im.CreateAt,
	}
}

type IPUpdateReq struct {
	// 标题
	Title string `json:"title"`
	// ip
	IP string `json:"ip"`
}

func (iur *IPUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"title":     iur.Title,
		"ip":        iur.IP,
		"update_by": openID,
		"update_at": time.Now(),
	}
}

// ip local resp
type City struct {
	City struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Continent struct {
		Code      string            `maxminddb:"code"`
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Location struct {
		AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
		Latitude       float64 `maxminddb:"latitude"`
		Longitude      float64 `maxminddb:"longitude"`
		MetroCode      uint    `maxminddb:"metro_code"`
		TimeZone       string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
	Postal struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
		Type              string            `maxminddb:"type"`
	} `maxminddb:"represented_country"`
	Subdivisions []struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		IsoCode   string            `maxminddb:"iso_code"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
	Traits struct {
		IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy"`
		IsSatelliteProvider bool `maxminddb:"is_satellite_provider"`
	} `maxminddb:"traits"`
}
