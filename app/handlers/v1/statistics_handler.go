package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/service"
	"js_statistics/commom/tools"
	"js_statistics/constant"

	"github.com/kataras/iris/v12"
)

type StatisticHandler struct {
	handlers.BaseHandler
	WIPSvc service.WhiteIPService
	Svc    service.StcService
}

func NewStatisticHandler() *StatisticHandler {
	return &StatisticHandler{
		WIPSvc: service.GetWhiteIPService(),
		Svc:    service.GetStcService(),
	}
}

// Create godoc
// @Summary js链接(开发中, 暂不可调)
// @Description js链接
// @Tags JS链接
// @Param sign path string true "js唯一字符串"
// @Router / [get]
func (sh *StatisticHandler) FilterJS(ctx iris.Context) {
	ip := tools.GetRemoteAddr(ctx)
	/*
		1. 是否是ip //TODO 后期域名的处理(目前不清楚获取的ip是否有域名)
		2. ip/
		//TODO cdn白名单先放着
		3. 判断地区屏蔽
		5. js跳转次数(记录存放次数，对比)
		6. 脚本的封禁小时(记录保存时间)
		4. 判断是pc端、移动端(安卓、ios)
		7. 来源判断(无，关键词、搜索引擎 判断)
		8. 跳转管理的地址(保证开关是开启可用)
		9. 跳转方式 --> 输出跳转代码
		10. 条件不满足 --> 伪装内容设置 --> 空白页
	*/
	// domain := ctx.Request().Host
	// agent := ctx.Request().UserAgent()

	sign := ctx.Params().Get("sign")
	agent := ctx.Request().UserAgent()
	origin := ctx.GetHeader("Origin")
	if len(sign) == 0 && len(agent) == 0 && len(origin) == 0 {
		tools.BeyondRuleRedirect(ctx)
	}
	visitType, cookie := tools.GetVisitType(ctx)
	// 看是否在ip白名单中
	if sh.WIPSvc.IsExistByIP(ip) {
		URL, ex := sh.Svc.Process(sign, agent, origin, ip, cookie, visitType)
		if ex != nil && len(URL) == 0 {
			tools.BeyondRuleRedirect(ctx)
		}
		ctx.WriteString(URL)
	} else {
		if !tools.IsValidIP(ip) {
			// 非法ip
			tools.BeyondRuleRedirect(ctx)
		}
		//
		if !IsValidLocation(ip) {
			// 不是国内的ip 直接跳转空白页
			tools.BeyondRuleRedirect(ctx)
		}
		// TODO 为进行cdn白名单过滤
	}
}

func IsValidLocation(ip string) bool {
	location, ex := tools.OriginIPLocation(ip)
	if ex != nil {
		return false
	}
	return location.Country.IsoCode == constant.CN_ISO_CODE
}
