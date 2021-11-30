package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/middlewares"
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type RmHandler struct {
	handlers.BaseHandler
	Svc   service.RmService
	RlSvc service.RlService
}

func NewRedirectManageHandler() *RmHandler {
	return &RmHandler{
		Svc:   service.GetRmService(),
		RlSvc: service.GetRlService(),
	}
}

// Create godoc
// @Summary 创建跳转管理
// @Description 创建跳转管理管理
// @Tags 应用管理 - 跳转管理
// @Param parameters body vo.RedirectManageReq true "RedirectManageReq"
// @Success 200  "创建跳转管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirect [post]
func (jmh *RmHandler) Create(ctx iris.Context) mvc.Result {
	cdn := &vo.RedirectManageReq{}
	if err := ctx.ReadJSON(cdn); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := jmh.Svc.Create(jmh.UserName, cdn)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询跳转管理列表
// @Description 查询跳转管理列表
// @Tags 应用管理 - 跳转管理
// @Param cid path string true "js分类id"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.RedirectManageResp} "查询跳转管理列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirects/category/{cid} [get]
func (jmh *RmHandler) ListByCategoryID(ctx iris.Context) mvc.Result {
	cid, err := ctx.Params().GetInt64(constant.CategoryID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := jmh.Svc.ListByCategoryID(params, cid)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询跳转管理
// @Description 查询跳转管理信息
// @Tags 应用管理 - 跳转管理
// @Param id path string true "跳转管理id"
// @Success 200 {object} vo.RedirectManageResp "查询域名成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirect/{id} [get]
func (jmh *RmHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := jmh.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改跳转管理
// @Description 修改跳转管理信息
// @Tags 应用管理 - 跳转管理
// @Param id path string true "跳转管理id"
// @Param parameters body vo.RedirectManageUpdateReq true "RedirectManageUpdateReq"
// @Success 200 "修改跳转管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirect/{id} [put]
func (jmh *RmHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.RedirectManageUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := jmh.Svc.Update(jmh.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除跳转管理
// @Description 删除跳转管理信息
// @Tags 应用管理 - 跳转管理
// @Param id path string true "跳转管理id"
// @Success 200 "删除跳转管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirect/{id} [delete]
func (jmh *RmHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := jmh.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量删除跳转管理
// @Description 批量删除跳转管理信息
// @Tags 应用管理 - 跳转管理
// @Param ids query string true "跳转管理ids, `,` 连接"
// @Success 200 "批量删除跳转管理管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirect/multi [delete]
func (jmh *RmHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := jmh.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询跳转管理日志列表
// @Description 查询跳转管理日志列表
// @Tags 应用管理 - 跳转管理 - 操作日志
// @Param cid path string true "js分类id"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Success 200 {object} vo.DataPagination{data=[]vo.RedirectLogResp} "查询跳转管理日志列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/redirect/logs/category/{cid} [get]
func (jmh *RmHandler) ListLogByCategoryID(ctx iris.Context) mvc.Result {
	cid, err := ctx.Params().GetInt64(constant.CategoryID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := jmh.RlSvc.ListByCategoryID(params, cid)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (jmh *RmHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/redirect", "Create", middlewares.RecordSystemLog("Create", "", "创建跳转信息成功"))
	b.Handle(iris.MethodGet, "/redirects/category/{cid:string}", "ListByCategoryID")
	b.Handle(iris.MethodGet, "/redirect/logs/category/{cid:string}", "ListLogByCategoryID")
	b.Handle(iris.MethodGet, "/redirect/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/redirect/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新跳转信息成功"))
	b.Handle(iris.MethodDelete, "/redirect/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除跳转信息成功"))
	b.Handle(iris.MethodDelete, "/redirect/multi", "MultiDelete", middlewares.RecordSystemLog("MultiDelete", "ids", "批量删除跳转信息成功"))
}
