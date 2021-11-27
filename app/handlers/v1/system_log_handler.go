package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type SyslogHandler struct {
	handlers.BaseHandler
	Svc service.SyslogService
}

func NewSyslogHandler() *SyslogHandler {
	return &SyslogHandler{
		Svc: service.GetSyslogService(),
	}
}

// Create godoc
// @Summary 查询操作日志
// @Description 查询操作日志列表
// @Tags 权限管理 - 操作日志
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.SystemLogResp} "查询操作日志列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/systemlogs [get]
func (slh *SyslogHandler) List(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := slh.Svc.List(params)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (slh *SyslogHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/systemlogs", "List")
}
