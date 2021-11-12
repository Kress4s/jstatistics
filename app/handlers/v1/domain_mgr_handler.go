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

type DomainHandler struct {
	handlers.BaseHandler
	Svc service.DomainService
}

func NewDomainHandler() *DomainHandler {
	return &DomainHandler{
		Svc: service.GetDomainService(),
	}
}

// Create godoc
// @Summary 创建域名
// @Description 创建域名管理
// @Tags 应用管理 - 域名管理
// @Param parameters body vo.DomainReq true "DomainReq"
// @Success 200  "创建域名成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/domain [post]
func (dh *DomainHandler) Create(ctx iris.Context) mvc.Result {
	domain := &vo.DomainReq{}
	if err := ctx.ReadJSON(domain); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := dh.Svc.Create(dh.UserName, domain)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询域名列表
// @Description 查询域名管理列表
// @Tags 应用管理 - 域名管理
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.DomainResp} "查询域名列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/domains [get]
func (dh *DomainHandler) List(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := dh.Svc.List(params)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询域名
// @Description 查询域名信息
// @Tags 应用管理 - 域名管理
// @Param id path string true "域名管理id"
// @Success 200 {object} vo.DomainResp "查询域名成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/domain/{id} [get]
func (dh *DomainHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetUint(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := dh.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改域名管理
// @Description 修改域名管理信息
// @Tags 应用管理 - 域名管理
// @Param id path string true "域名管理id"
// @Param parameters body vo.DomainUpdateReq true "DomainUpdateReq"
// @Success 200 "修改域名成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/domain/{id} [put]
func (dh *DomainHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetUint(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	user := &vo.DomainUpdateReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := dh.Svc.Update(dh.UserName, id, user)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除域名管理
// @Description 删除域名管理信息
// @Tags 应用管理 - 域名管理
// @Param id path string true "域名管理id"
// @Success 200 "删除域名成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/domain/{id} [delete]
func (dh *DomainHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetUint(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := dh.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量删除域名管理
// @Description 批量删除域名管理信息
// @Tags 应用管理 - 域名管理
// @Param ids query string true "域名管理ids, `,`连接"
// @Success 200 "批量删除域名成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/domain/multi [delete]
func (dh *DomainHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := dh.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (dh *DomainHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/domain", "Create")
	b.Handle(iris.MethodGet, "/domains", "List")
	b.Handle(iris.MethodGet, "/domain/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/domain/{id:string}", "Update")
	b.Handle(iris.MethodDelete, "/domain/{id:string}", "Delete")
	b.Handle(iris.MethodDelete, "/domain/multi", "MultiDelete")
}
