package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/middlewares"
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/commom/tools"
	"js_statistics/constant"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type BlackIPHandler struct {
	handlers.BaseHandler
	Svc service.BlackIPService
}

func NewBlackIPHandler() *BlackIPHandler {
	return &BlackIPHandler{
		Svc: service.GetBlackIPService(),
	}
}

// Create godoc
// @Summary 创建ip库
// @Description 创建ip库管理
// @Tags 应用管理 - ip库管理
// @Param parameters body vo.BlackIPReq true "BlackIPReq"
// @Success 200  "创建ip库成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ip [post]
func (ih *BlackIPHandler) Create(ctx iris.Context) mvc.Result {
	ip := &vo.BlackIPReq{}
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
// @Summary 查询ip库列表
// @Description 查询ip库管理列表
// @Tags 应用管理 - ip库管理
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.BlackIPResp} "查询ip列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ips [get]
func (ih *BlackIPHandler) List(ctx iris.Context) mvc.Result {
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
// @Summary 查询ip
// @Description 查询ip库信息
// @Tags 应用管理 - ip库管理
// @Param id path string true "ip管理id"
// @Success 200 {object} vo.BlackIPResp "查询ip成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ip/{id} [get]
func (ih *BlackIPHandler) Get(ctx iris.Context) mvc.Result {
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
// @Summary 修改ip管理
// @Description 修改ip库管理信息
// @Tags 应用管理 - ip库管理
// @Param id path string true "ip管理库id"
// @Param parameters body vo.BlackIPUpdateReq true "BlackIPUpdateReq"
// @Success 200 "修改ip成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ip/{id} [put]
func (ih *BlackIPHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	user := &vo.BlackIPUpdateReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ih.Svc.Update(ih.UserName, id, user)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除ip库管理
// @Description 删除ip库管理信息
// @Tags 应用管理 - ip库管理
// @Param id path string true "ip管理id"
// @Success 200 "删除ip成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ip/{id} [delete]
func (ih *BlackIPHandler) Delete(ctx iris.Context) mvc.Result {
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
// @Summary 批量删除ip库
// @Description 批量删除ip库信息
// @Tags 应用管理 - ip库管理
// @Param ids query string true "ip库管理ids, `,` 连接"
// @Success 200 "批量删除ip库管理成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ip/multi [delete]
func (ih *BlackIPHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := ih.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询ip所在库(归属地)
// @Description 查询ip所在库(归属地)
// @Tags 应用管理 - ip库管理
// @Param ip query string true "ip"
// @Success 200 {object} tools.Location "查询ip所在库(归属地)成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/black/ip/where [get]
func (ih *BlackIPHandler) IPLocationSearch(ctx iris.Context) mvc.Result {
	ip := ctx.URLParam(constant.IP)
	if len(ip) == 0 {
		return response.Error(exception.New(response.ExceptionInvalidRequestParameters, "ip不能为空"))
	}
	localtion, ex := tools.LocationIP(ip)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(localtion)
}

// BeforeActivation 初始化路由
func (ih *BlackIPHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/black/ip", "Create", middlewares.RecordSystemLog("Create", "", "创建黑名单成功"))
	b.Handle(iris.MethodGet, "/black/ips", "List")
	b.Handle(iris.MethodGet, "/black/ip/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/black/ip/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新黑名单成功"))
	b.Handle(iris.MethodDelete, "/black/ip/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除黑名单成功"))
	b.Handle(iris.MethodDelete, "/black/ip/multi", "MultiDelete", middlewares.RecordSystemLog("MultiDelete", "ids", "批量删除黑名单成功"))
	b.Handle(iris.MethodGet, "/black/ip/where", "IPLocationSearch")
}
