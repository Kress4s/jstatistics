package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type FromAnalysisHandler struct {
	handlers.BaseHandler
	Svc service.FaService
}

func NewFromAnalysisHandler() *FromAnalysisHandler {
	return &FromAnalysisHandler{
		Svc: service.GetFaService(),
	}
}

// FromStatistic godoc
// @Summary 来路统计数据查询
// @Description 查询来路统计数据
// @Tags 数据统计 - 来路统计
// @Param begin_at query string true "时间格式: 2021-08-24"
// @Param end_at query string true "时间格式: 2021-08-31"
// @Param pid query int true "JS主分类ID"
// @Param cid query int false "JS分类ID"
// @Param jid query int false "JS ID"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.FromAnalysisResp} "查询来路统计数据"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/analysis/flow/from/statistic [get]
func (fah *FromAnalysisHandler) FromStatistic(ctx iris.Context) mvc.Result {
	param, ex := handlers.GetJSFilterParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	pageInfo, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	beginAt, endAt, ex := handlers.GetTimeScopeParam(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := fah.Svc.FromStatistic(param, pageInfo, beginAt, endAt)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (fah *FromAnalysisHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/flow/from/statistic", "FromStatistic")
}
