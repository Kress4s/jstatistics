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

type RoleHandler struct {
	handlers.BaseHandler
	Svc service.RoleService
}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{
		Svc: service.GetRoleService(),
	}
}

// Create godoc
// @Summary 创建角色
// @Description 创建角色
// @Tags 权限管理 - 管理组
// @Param parameters body vo.RoleReq true "RoleReq"
// @Success 200  "创建角色成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/role [post]
func (rh *RoleHandler) Create(ctx iris.Context) mvc.Result {
	param := &vo.RoleReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := rh.Svc.Create(rh.UserName, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询角色
// @Description 查询角色信息
// @Tags 权限管理 - 管理组
// @Param id path string true "角色id"
// @Success 200 {object} vo.RoleResp "查询角色成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/role/{id} [get]
func (rh *RoleHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	resp, ex := rh.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 查询角色列表
// @Description 查询角色列表信息
// @Tags 权限管理 - 管理组
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.RoleResp} "查询角色列表成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/roles [get]
func (rh *RoleHandler) List(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	resp, ex := rh.Svc.List(params)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改角色
// @Description 修改角色信息
// @Tags 权限管理 - 管理组
// @Param id path string true "角色id"
// @Param parameters body vo.RoleUpdateReq true "RoleUpdateReq"
// @Success 200 "修改角色成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/role/{id} [put]
func (rh *RoleHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.RoleUpdateReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := rh.Svc.Update(rh.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除角色
// @Description 删除角色
// @Tags 权限管理 - 管理组
// @Param id path string true "角色id"
// @Success 200 "删除角色成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/role/{id} [delete]
func (rh *RoleHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := rh.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (rh *RoleHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/role", "Create")
	b.Handle(iris.MethodGet, "/role/{id:string}", "Get")
	b.Handle(iris.MethodGet, "/roles", "List")
	b.Handle(iris.MethodPut, "/role/{id:string}", "Update")
	b.Handle(iris.MethodDelete, "/role/{id:string}", "Delete")
}
