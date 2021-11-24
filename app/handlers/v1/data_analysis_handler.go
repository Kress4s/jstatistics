package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DataAnalysisHandler struct {
	handlers.BaseHandler
	Svc service.DaService
}

func NewDataAnalysisHandler() *DataAnalysisHandler {
	return &DataAnalysisHandler{
		Svc: service.GetDaService(),
	}
}

// TodayIP godoc
// @Summary 查询今日IP统计数
// @Description 查询今日IP统计数
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.TodayIP "查询今日IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/today_ip [get]
func (dah *DataAnalysisHandler) TodayIP(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.TodayIP(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// ThisMonthIP godoc
// @Summary 查询昨日IP统计数
// @Description 查询昨日IP统计数
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.YesterdayIP "查询昨日IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/yesterday_ip [get]
func (dah *DataAnalysisHandler) YesterdayIP(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.YesterdayIP(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// ThisMonthIP godoc
// @Summary 查询本月IP统计数
// @Description 查询本月IP统计数
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.ThisMonthIP "查询本月IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/this_month_ip [get]
func (dah *DataAnalysisHandler) ThisMonthIP(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.ThisMonthIP(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// LastMonthIP godoc
// @Summary 查询上月IP统计数
// @Description 查询上月IP统计数
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.LastMonthIP "查询上月IP统计数成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/last_month_ip [get]
func (dah *DataAnalysisHandler) LastMonthIP(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.LastMonthIP(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// IPAndUVisit godoc
// @Summary 查询时间范围内IP、UV的统计(过去七天、过去30天、自选时间范围)
// @Description 查询时间范围内IP、UV的统计
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Param begin_at query string true "时间格式: 2021-08-24"
// @Param end_at query string true "时间格式: 2021-08-31"
// @Success 200 {object} vo.HomeIPAndUVisit "查询时间范围内IP、UV的统计成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/timescope/ip_uv [get]
func (dah *DataAnalysisHandler) IPAndUVisit(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	beginAt, endAt, ex := handlers.GetTimeScopeParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.IPAndUVisit(param, beginAt, endAt)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// TodayIPAndUVisit godoc
// @Summary 查询今日IP、UV的统计
// @Description 查询今日IP、UV的统计
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.HomeIPAndUVisit "查询今日IP、UV的统计成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/today/ip_uv [get]
func (dah *DataAnalysisHandler) TodayIPAndUVisit(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.TodayIPAndUVisit(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// YesterdayIPAndUVisit godoc
// @Summary 查询昨日IP、UV的统计
// @Description 查询昨日IP、UV的统计
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.HomeIPAndUVisit "查询昨日IP、UV的统计成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/yesterday/ip_uv [get]
func (dah *DataAnalysisHandler) YesterdayIPAndUVisit(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.YesterdayIPAndUVisit(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// FromNowIPAndUVisit godoc
// @Summary 查询 开始至今 IP、UV的统计
// @Description 查询开始至今IP、UV的统计
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Success 200 {object} vo.HomeIPAndUVisit "查询昨日IP、UV的统计成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/from_now/ip_uv [get]
func (dah *DataAnalysisHandler) FromNowIPAndUVisit(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.FromNowIPAndUVisit(param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// TodayFlowData godoc
// @Summary 查询今日流量数据
// @Description 查询今日流量数据
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.FlowDataResp} "查询今日流量数据成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/today/flowdata [get]
func (dah *DataAnalysisHandler) TodayFlowData(ctx iris.Context) mvc.Result {
	pageInfo, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.TodayFlowData(param, pageInfo)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// YesterdayFlowData godoc
// @Summary 查询昨日流量数据
// @Description 查询昨日流量数据
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.FlowDataResp} "查询昨日流量数据成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/yesterday/flowdata [get]
func (dah *DataAnalysisHandler) YesterdayFlowData(ctx iris.Context) mvc.Result {
	pageInfo, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.YesterdayFlowData(param, pageInfo)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// TimeScopeFlowData godoc
// @Summary 查询时间范围流量数据(过去七天、过去30天、自选时间范围)
// @Description 查询时间范围内流量数据
// @Tags 数据统计 - 流量统计
// @Param begin_at query string true "时间格式: 2021-08-24"
// @Param end_at query string true "时间格式: 2021-08-31"
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.FlowDataResp} "查询开始时间范围流量数据成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/timescope/flowdata [get]
func (dah *DataAnalysisHandler) TimeScopeFlowData(ctx iris.Context) mvc.Result {
	beginAt, endAt, ex := handlers.GetTimeScopeParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	pageInfo, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.TimeScopeFlowData(param, pageInfo, beginAt, endAt)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// FromNowFlowData godoc
// @Summary 查询开始至今流量数据
// @Description 查询开始至今流量数据
// @Tags 数据统计 - 流量统计
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.FlowDataResp} "查询开始时间范围流量数据成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/from_now/flowdata [get]
func (dah *DataAnalysisHandler) FromNowFlowData(ctx iris.Context) mvc.Result {
	pageInfo, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dah.Svc.FromNowFlowData(param, pageInfo)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (dah *DataAnalysisHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/flow/today_ip", "TodayIP")
	b.Handle(iris.MethodGet, "/flow/yesterday_ip", "YesterdayIP")
	b.Handle(iris.MethodGet, "/flow/this_month_ip", "ThisMonthIP")
	b.Handle(iris.MethodGet, "/flow/last_month_ip", "LastMonthIP")
	b.Handle(iris.MethodGet, "/flow/timescope/ip_uv", "IPAndUVisit")
	b.Handle(iris.MethodGet, "/flow/today/ip_uv", "TodayIPAndUVisit")
	b.Handle(iris.MethodGet, "/flow/yesterday/ip_uv", "YesterdayIPAndUVisit")
	b.Handle(iris.MethodGet, "/flow/from_now/ip_uv", "FromNowIPAndUVisit")
	b.Handle(iris.MethodGet, "/flow/today/flowdata", "TodayFlowData")
	b.Handle(iris.MethodGet, "/flow/yesterday/flowdata", "YesterdayFlowData")
	b.Handle(iris.MethodGet, "/flow/timescope/flowdata", "TimeScopeFlowData")
	b.Handle(iris.MethodGet, "/flow/from_now/flowdata", "FromNowFlowData")
}
