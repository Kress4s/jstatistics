package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type JsmHandler struct {
	handlers.BaseHandler
	Svc service.JsmService
}

func NewJsManageHandler() *JsmHandler {
	return &JsmHandler{
		Svc: service.GetJsmService(),
	}
}

// Create godoc
// @Summary 创建js管理
// @Description 创建js管理管理
// @Tags 应用管理 - js管理
// @Param parameters body vo.JsManageReq true "JsManageReq"
// @Success 200  "创建js管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_manage [post]
func (jmh *JsmHandler) Create(ctx iris.Context) mvc.Result {
	cdn := &vo.JsManageReq{}
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
// @Summary 查询js管理列表通过js分类ID
// @Description 查询js管理列表通过js分类ID
// @Tags 应用管理 - js管理
// @Param cid path string true "js分类id"
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.JsManageResp} "查询js管理列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_categories/category/{cid} [get]
func (jmh *JsmHandler) ListByCategoryID(ctx iris.Context) mvc.Result {
	pid, err := ctx.Params().GetInt64(constant.CategoryID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := jmh.Svc.ListByCategoryID(params, pid)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询js管理
// @Description 查询js管理信息
// @Tags 应用管理 - js管理
// @Param id path string true "js管理id"
// @Success 200 {object} vo.JsManageResp "查询js管理成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_manage/{id} [get]
func (jmh *JsmHandler) Get(ctx iris.Context) mvc.Result {
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
// @Summary 修改js管理
// @Description 修改js管理信息
// @Tags 应用管理 - js管理
// @Param id path string true "js管理id"
// @Param parameters body vo.JsManageUpdateReq true "JsManageUpdateReq"
// @Success 200 "修改js管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_manage/{id} [put]
func (jmh *JsmHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.JsManageUpdateReq{}
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
// @Summary 删除js管理
// @Description 删除js管理信息
// @Tags 应用管理 - js管理
// @Param id path string true "js管理id"
// @Success 200 "删除js管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_manage/{id} [delete]
func (jmh *JsmHandler) Delete(ctx iris.Context) mvc.Result {
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
// @Summary 批量删除js管理
// @Description 批量删除js管理信息
// @Tags 应用管理 - js管理
// @Param ids query string true "js管理ids, `,` 连接"
// @Success 200 "批量删除js管理管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_manage/multi [delete]
func (jmh *JsmHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := jmh.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询js地址
// @Description 查询js地址信息
// @Tags 应用管理 - js管理
// @Param id path string true "js管理id"
// @Success 200 {object} vo.JSiteResp "查询js地址成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_manage/js_site/{id} [get]
func (jmh *JsmHandler) GetJSiteByID(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := jmh.Svc.GetJSiteByID(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// BeforeActivation 初始化路由
func (jmh *JsmHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/js_manage", "Create")
	b.Handle(iris.MethodGet, "/js_categories/category/{cid:string}", "ListByCategoryID")
	b.Handle(iris.MethodGet, "/js_manage/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/js_manage/{id:string}", "Update")
	b.Handle(iris.MethodDelete, "/js_manage/{id:string}", "Delete")
	b.Handle(iris.MethodDelete, "/js_manage/multi", "MultiDelete")
	b.Handle(iris.MethodGet, "/js_manage/js_site/{id:string}", "GetJSiteByID")
}
