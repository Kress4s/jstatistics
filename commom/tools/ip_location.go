package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"js_statistics/app/response"
	"js_statistics/constant"
	"js_statistics/exception"
	"log"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

type Location struct {
	// IP       string `json:"ip"`
	// Province string `json:"pro"`
	// City     string `json:"city"`
	// Address  string `json:"addr"`
	// 国家
	Country string `json:"country"`
	// 省份
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
}

func IPLocation(ip string) (*Location, exception.Exception) {
	url := fmt.Sprintf(constant.IPLocation, ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionHttpRequestError, err)
	}
	defer resp.Body.Close()
	localtion := &Location{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, localtion)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionUnmarshalJSON, err)
	}
	return localtion, nil
}

func LocationIP(ip string) (*Location, exception.Exception) {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	_ip := net.ParseIP(ip)
	record, err := db.City(_ip)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionPraseIPLocationError, err)
	}
	localtion := &Location{}
	localtion.Country = record.Country.Names["zh-CN"]
	if len(record.Subdivisions) != 0 {
		localtion.Province = record.Subdivisions[0].Names["zh-CN"]
	}
	if &record.City != nil {
		localtion.City = record.City.Names["zh-CN"]
	}
	return localtion, nil
}
