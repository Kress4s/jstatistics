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

type PermissionHandler struct {
	handlers.BaseHandler
	Svc service.PermissionService
}

func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{
		Svc: service.GetPermissionService(),
	}
}

// Create godoc
// @Summary 创建权限规则
// @Description 创建权限规则
// @Tags 权限管理 - 权限规则
// @Param parameters body vo.PermissionReq true "PermissionReq"
// @Success 200  "创建权限规则成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/rule [post]
func (ph *PermissionHandler) Create(ctx iris.Context) mvc.Result {
	permission := &vo.PermissionReq{}
	if err := ctx.ReadJSON(permission); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ph.Svc.Create(ph.UserName, permission)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 获取权限规则
// @Description 获取权限规则
// @Tags 权限管理 - 权限规则
// @Param id path string true "权限规则id"
// @Success 200 {object} vo.PermissionResp "查询权限规则成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/rule/{id} [get]
func (ph *PermissionHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	res, ex := ph.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// GetAll godoc
// @Summary 获取全部权限规则
// @Description 获取全部权限规则
// @Tags 权限管理 - 权限规则
// @Success 200 {array} vo.PermissionResp "查询权限规则成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/rules/all [get]
func (ph *PermissionHandler) GetAll(ctx iris.Context) mvc.Result {
	res, ex := ph.Svc.GetAll()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// Create godoc
// @Summary 获取权限规则树
// @Description 获取权限规则树
// @Tags 权限管理 - 权限规则
// @Success 200 {object} vo.PermissionTree "查询权限规则成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/rules [get]
func (ph *PermissionHandler) GetPermissionTree(ctx iris.Context) mvc.Result {
	res, ex := ph.Svc.GetPermissionTree()
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// Create godoc
// @Summary 修改权限规则
// @Description 修改权限规则
// @Tags 权限管理 - 权限规则
// @Param id path string true "权限规则id"
// @Param parameters body vo.PermissionUpdateReq true "PermissionUpdateReq"
// @Success 200  "修改权限规则成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/rule/{id} [put]
func (ph *PermissionHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	p := &vo.PermissionUpdateReq{}
	if err := ctx.ReadJSON(p); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := ph.Svc.Update(ph.UserName, id, p)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除权限规则
// @Description 删除权限规则
// @Tags 权限管理 - 权限规则
// @Param id path string true "权限规则id"
// @Success 200  "删除权限规则成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/rule/{id} [delete]
func (ph *PermissionHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := ph.Svc.Delete(ph.UserName, id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ph *PermissionHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/rule", "Create", middlewares.RecordSystemLog("Create", "", "创建权限规则成功"))
	b.Handle(iris.MethodGet, "/rule/{id:string}", "Get")
	b.Handle(iris.MethodGet, "/rules/all", "GetAll")
	b.Handle(iris.MethodGet, "/rules", "GetPermissionTree")
	b.Handle(iris.MethodPut, "/rule/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新权限规则成功"))
	b.Handle(iris.MethodDelete, "/rule/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除权限规则成功"))
}
