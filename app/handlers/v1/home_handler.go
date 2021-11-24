package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type HomeHandler struct {
	handlers.BaseHandler
	Svc service.HomeService
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{
		Svc: service.GetHomeService(),
	}
}

// TodayIP godoc
// @Summary 查询今日IP
// @Description 查询今日IP统计数
// @Tags 主页 - 统计
// @Success 200 {object} vo.TodayIP "查询今日IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/today_ip [get]
func (hh *HomeHandler) TodayIP(ctx iris.Context) mvc.Result {
	resp, ex := hh.Svc.TodayIP()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// YesterdayIP godoc
// @Summary 查询昨日IP
// @Description 查询昨日IP统计数
// @Tags 主页 - 统计
// @Success 200 {object} vo.YesterdayIP "查询昨日IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/yesterday_ip [get]
func (hh *HomeHandler) YesterdayIP(ctx iris.Context) mvc.Result {
	resp, ex := hh.Svc.YesterdayIP()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// ThisMonthIP godoc
// @Summary 查询本月IP
// @Description 查询本月IP统计数
// @Tags 主页 - 统计
// @Success 200 {object} vo.ThisMonthIP "查询本月IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/this_month_ip [get]
func (hh *HomeHandler) ThisMonthIP(ctx iris.Context) mvc.Result {
	resp, ex := hh.Svc.ThisMonthIP()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// LastMonthIP godoc
// @Summary 查询上月IP
// @Description 查询上月IP统计数
// @Tags 主页 - 统计
// @Success 200 {object} vo.LastMonthIP "查询上月IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/last_month_ip [get]
func (hh *HomeHandler) LastMonthIP(ctx iris.Context) mvc.Result {
	resp, ex := hh.Svc.LastMonthIP()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// LastMonthIP godoc
// @Summary 查询IP、UV访问量统计
// @Description 访问量统计
// @Tags 主页 - 统计
// @Success 200 {object} vo.HomeIPAndUVisit "查询访问量统计成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/ip_uv/statistic [get]
func (hh *HomeHandler) IPAndUVisit(ctx iris.Context) mvc.Result {
	resp, ex := hh.Svc.IPAndUVisit()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// LastMonthIP godoc
// @Summary 查询IP地域地图统计
// @Description IP地域地图统计详情
// @Tags 主页 - 统计
// @Success 200 {array} vo.RegionStatisticResp "查询访问量统计成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/region/statistic [get]
func (hh *HomeHandler) RegionStatistic(ctx iris.Context) mvc.Result {
	resp, ex := hh.Svc.RegionStatistic()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// JSVisitStatistic godoc
// @Summary 查询js流量排行榜
// @Description js流量排行榜详情
// @Tags 主页 - 统计
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination "查询js流量排行榜成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/home/js/rank [get]
func (hh *HomeHandler) JSVisitStatistic(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := hh.Svc.JSVisitStatistic(params)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (hh *HomeHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/today_ip", "TodayIP")
	b.Handle(iris.MethodGet, "/yesterday_ip", "YesterdayIP")
	b.Handle(iris.MethodGet, "/this_month_ip", "ThisMonthIP")
	b.Handle(iris.MethodGet, "/last_month_ip", "LastMonthIP")
	b.Handle(iris.MethodGet, "/ip_uv/statistic", "IPAndUVisit")
	b.Handle(iris.MethodGet, "/region/statistic", "RegionStatistic")
	b.Handle(iris.MethodGet, "/js/rank", "JSVisitStatistic")
}
