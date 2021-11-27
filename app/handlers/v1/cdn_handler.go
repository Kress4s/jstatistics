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

type CdnHandler struct {
	handlers.BaseHandler
	Svc service.CdnService
}

func NewCdnHandler() *CdnHandler {
	return &CdnHandler{
		Svc: service.GetCdnService(),
	}
}

// Create godoc
// @Summary 创建cdn白名单
// @Description 创建cdn白名单管理
// @Tags 应用管理 - cdn白名单
// @Param parameters body vo.CDNReq true "CDNReq"
// @Success 200  "创建cdn白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/cdn [post]
func (ch *CdnHandler) Create(ctx iris.Context) mvc.Result {
	cdn := &vo.CDNReq{}
	if err := ctx.ReadJSON(cdn); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ch.Svc.Create(ch.UserName, cdn)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询cdn白名单列表
// @Description 查询cdn白名单列表
// @Tags 应用管理 - cdn白名单
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.CDNResp} "查询cdn白名单列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/cdns [get]
func (ch *CdnHandler) List(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := ch.Svc.List(params)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询cdn白名单
// @Description 查询cdn白名单信息
// @Tags 应用管理 - cdn白名单
// @Param id path string true "cdn白名单id"
// @Success 200 {object} vo.CDNResp "查询域名成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/cdn/{id} [get]
func (ch *CdnHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := ch.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改cdn白名单
// @Description 修改cdn白名单信息
// @Tags 应用管理 - cdn白名单
// @Param id path string true "cdn白名单id"
// @Param parameters body vo.CDNUpdateReq true "CDNUpdateReq"
// @Success 200 "修改cdn白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/cdn/{id} [put]
func (ch *CdnHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.CDNUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ch.Svc.Update(ch.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除cdn白名单
// @Description 删除cdn白名单信息
// @Tags 应用管理 - cdn白名单
// @Param id path string true "cdn白名单id"
// @Success 200 "删除cdn白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/cdn/{id} [delete]
func (ch *CdnHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := ch.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量删除cdn白名单
// @Description 批量删除cdn白名单信息
// @Tags 应用管理 - cdn白名单
// @Param ids query string true "cdn白名单ids, `,` 连接"
// @Success 200 "批量删除cdn白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/cdn/multi [delete]
func (ch *CdnHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := ch.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ch *CdnHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/cdn", "Create", middlewares.RecordSystemLog("Create", "", "创建cdn成功"))
	b.Handle(iris.MethodGet, "/cdns", "List")
	b.Handle(iris.MethodGet, "/cdn/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/cdn/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新cdn成功"))
	b.Handle(iris.MethodDelete, "/cdn/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除cdn成功"))
	b.Handle(iris.MethodDelete, "/cdn/multi", "MultiDelete", middlewares.RecordSystemLog("MultiDelete", "ids", "批量删除cdn成功"))
}
