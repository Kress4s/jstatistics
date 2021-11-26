package tools

import (
	"js_statistics/constant"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

func IsInRuleClient(client int64, jsClient []int64) bool {
	isValid := false
	for i := range jsClient {
		if jsClient[i] == client {
			isValid = !isValid
			break
		}
	}
	return isValid
}

func GetClintType(agent string) int {
	if strings.ContainsAny(strings.ToLower(agent), constant.MOBILE) {
		return constant.MobileType
	}
	return constant.PCType
}

func GetDeviceType(agent string) int {
	switch {
	case strings.ContainsAny(strings.ToLower(agent), constant.IOS):
		return constant.IOSRedirectType
	case strings.ContainsAny(strings.ToLower(agent), constant.Android):
		return constant.AndroidRedirectType
	default:
		return constant.PCRedirectType
	}
}

func GetEngineType(agent string) (bool, int64) {
	if strings.Contains(strings.ToLower(agent), constant.BaiduSearch) {
		return true, constant.Baidu
	}
	if strings.Contains(strings.ToLower(agent), constant.UCSearch) {
		return true, constant.UC
	}
	if strings.Contains(strings.ToLower(agent), constant.SLLSearch) {
		return true, constant.SLL
	}
	if strings.Contains(strings.ToLower(agent), constant.SOU_GOUSearch) {
		return true, constant.SOU_GOU
	}
	if strings.ContainsAny(strings.ToLower(agent), constant.GOOGLESearch) {
		return true, constant.GOOGLE
	}
	if strings.Contains(strings.ToLower(agent), constant.BingSearch) {
		return true, constant.Bing
	}
	return false, -1
}

func GetRedirectCode(rt int) string {
	return ""
}

func GetVisitType(ctx iris.Context) (int, string) {
	cookie := ctx.GetCookie(constant.CookieKey)
	if len(cookie) == 0 {
		ctx.SetCookie(&http.Cookie{
			Name:    constant.CookieKey,
			Value:   RandCookie(15),
			Path:    "/",
			Expires: time.Now().Add(24 * time.Hour),
		})
		return constant.IPVisit, ""
	}
	return constant.UVVisit, cookie
}
