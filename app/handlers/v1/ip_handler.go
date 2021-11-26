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

type IPHandler struct {
	handlers.BaseHandler
	Svc service.WhiteIPService
}

func NewIPHandler() *IPHandler {
	return &IPHandler{
		Svc: service.GetWhiteIPService(),
	}
}

// Create godoc
// @Summary 创建ip白名单
// @Description 创建ip白名单
// @Tags 应用管理 - ip白名单
// @Param parameters body vo.IPReq true "IPReq"
// @Success 200  "创建ip白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/ip [post]
func (ih *IPHandler) Create(ctx iris.Context) mvc.Result {
	ip := &vo.IPReq{}
	if err := ctx.ReadJSON(ip); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ih.Svc.Create(ih.UserName, ip)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询ip白名单
// @Description 查询ip白名单列表
// @Tags 应用管理 - ip白名单
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.IPReq} "查询ip列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/ips [get]
func (ih *IPHandler) List(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := ih.Svc.List(params)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询ip白名单
// @Description 查询ip白名单信息
// @Tags 应用管理 - ip白名单
// @Param id path string true "ip白名单id"
// @Success 200 {object} vo.IPResp "查询ip白名单成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/ip/{id} [get]
func (ih *IPHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := ih.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改ip白名单
// @Description 修改ip白名单信息
// @Tags 应用管理 - ip白名单
// @Param id path string true "ip白名单id"
// @Param parameters body vo.IPUpdateReq true "IPUpdateReq"
// @Success 200 "修改ip白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/ip/{id} [put]
func (ih *IPHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ip := &vo.IPUpdateReq{}
	if err := ctx.ReadJSON(ip); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ih.Svc.Update(ih.UserName, id, ip)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除ip白名单
// @Description 删除ip白名单信息
// @Tags 应用管理 - ip白名单
// @Param id path string true "ip白名单id"
// @Success 200 "删除ip白名单成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/ip/{id} [delete]
func (ih *IPHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := ih.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 批量删除ip白名单
// @Description 批量删除ip白名单信息
// @Tags 应用管理 - ip白名单
// @Param ids query string true "ip白名单ids, `,` 连接"
// @Success 200 "批量删除ip白名单管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/ip/multi [delete]
func (ih *IPHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := ih.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ih *IPHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/ip", "Create")
	b.Handle(iris.MethodGet, "/ips", "List")
	b.Handle(iris.MethodGet, "/ip/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/ip/{id:string}", "Update")
	b.Handle(iris.MethodDelete, "/ip/{id:string}", "Delete")
	b.Handle(iris.MethodDelete, "/ip/multi", "MultiDelete")
}
