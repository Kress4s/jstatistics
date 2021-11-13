package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"js_statistics/app/response"
	"js_statistics/constant"
	"js_statistics/exception"
	"net/http"
)

type Location struct {
	IP       string `json:"ip"`
	Province string `json:"pro"`
	City     string `json:"city"`
	Address  string `json:"addr"`
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
