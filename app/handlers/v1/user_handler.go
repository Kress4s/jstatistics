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

type UserHandler struct {
	handlers.BaseHandler
	Svc service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		Svc: service.GetUserService(),
	}
}

// Create godoc
// @Summary 获取登录用户的信息
// @Description 登录用户信息
// @Tags 权限管理 - 管理员
// @Success 200 {object} vo.ProfileResp "获取用户信息成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/profile [get]
func (u *UserHandler) Profile(ctx iris.Context) mvc.Result {
	profile, ex := u.Svc.Profile(u.UserID)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(profile)
}

// Create godoc
// @Summary 创建用户
// @Description 创建用户
// @Tags 权限管理 - 管理员
// @Param parameters body vo.UserReq true "UserReq"
// @Success 200  "创建用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user [post]
func (u *UserHandler) Create(ctx iris.Context) mvc.Result {
	user := &vo.UserReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := u.Svc.Create(u.UserName, user)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询用户
// @Description 查询用户信息
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Success 200 {object} vo.ProfileResp "查询用户成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id} [get]
func (u *UserHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	res, ex := u.Svc.Get(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// Create godoc
// @Summary 查询用户列表
// @Description 查询用户列表信息
// @Tags 权限管理 - 管理员
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "用户昵称"
// @Param id query string false "id"
// @Success 200 {object} vo.DataPagination{data=[]vo.UserResp} "查询用户列表成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/users [get]
func (u *UserHandler) List(ctx iris.Context) mvc.Result {
	params, ex := handlers.GetPageInfo(ctx)
	if ex != nil {
		return response.Error(ex)
	}
	// id 过滤
	var id int64
	var err error
	if ctx.URLParamExists(constant.ID) {
		if ctx.URLParam(constant.ID) != "" {
			id, err = ctx.URLParamInt64(constant.ID)
			if err != nil {
				return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
			}
		}
	}
	resp, ex := u.Svc.List(params, id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// Create godoc
// @Summary 修改用户
// @Description 修改用户
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Param parameters body vo.UserUpdateReq true "UserUpdateReq"
// @Success 200 "修改用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id} [put]
func (u *UserHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	user := &vo.UserUpdateReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := u.Svc.Update(u.UserName, id, user)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 修改用户角色权限
// @Description 修改角色权限
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Param parameters body vo.UserUpdateRolesReq true "UserUpdateRolesReq"
// @Success 200 "修改用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id}/roles [put]
func (u *UserHandler) UpdateRoles(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.UserUpdateRolesReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := u.Svc.UpdateRoles(u.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 修改用户js分类权限
// @Description 修改用户js分类权限
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Param parameters body vo.UserUpdateJscAndJsReq true "UserUpdateJscAndJsReq"
// @Success 200 "修改用户分类权限成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id}/jsc_js [put]
func (u *UserHandler) UpdateJSCAndJS(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.UserUpdateJscAndJsReq{}
	if err := ctx.ReadJSON(param); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := u.Svc.UpdateJSCAndJS(u.UserName, id, param)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 删除用户
// @Description 删除用户
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Success 200 "删除用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id} [delete]
func (u *UserHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := u.Svc.Delete(u.UserName, id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 查询用户角色信息
// @Description 查询用户角色信息
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Success 200 {object} vo.RoleBriefResp "查询用户角色信息成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id}/roles [get]
func (u *UserHandler) GetRolesByUserID(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	res, ex := u.Svc.GetRolesByUserID(u.UserName, id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// GetCategoryByUserID godoc
// @Summary 查询用户js分类权限信息
// @Description 查询用户js分类权限信息
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Success 200 {object} vo.JsJscAndJsBriefResp "查询用户js分类信息成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id}/jsc_js [get]
func (u *UserHandler) GetJscAndJsByUserID(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	res, ex := u.Svc.GetJscAndJsByUserID(u.UserName, id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(res)
}

// Create godoc
// @Summary 查询(动态菜单)当前用户可展示的菜单权限
// @Description 查询当前用户可展示的菜单权限信息
// @Tags 权限管理 - 管理员
// @Success 200 {array}  vo.UserToMenusResp "动态菜单获取成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/menus [get]
func (u *UserHandler) GetUserMenus(ctx iris.Context) mvc.Result {
	resp, ex := u.Svc.GetUserMenus(u.UserID)
	if ex != nil {
		return response.Error(ex)
	}
	return response.JSON(resp)
}

// StatusChange godoc
// @Summary 修改用户状态
// @Description 修改用户状态信息
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Param status query bool true "用户修改的状态"
// @Success 200 "修改用户状态成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id} [patch]
func (u *UserHandler) StatusChange(ctx iris.Context) mvc.Result {
	status, err := ctx.URLParamBool(constant.Status)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	id, err := ctx.Params().GetInt64(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := u.Svc.StatusChange(u.UserName, id, status)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// MultiDelete godoc
// @Summary 批量删除用户
// @Description 批量删除用户信息
// @Tags 权限管理 - 管理员
// @Param ids query string true "用户ids, `,` 连接"
// @Success 200 "批量删除用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/multi [delete]
func (u *UserHandler) MultiDelete(ctx iris.Context) mvc.Result {
	ids := ctx.URLParam(constant.IDS)
	ex := u.Svc.MultiDelete(ids)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (u *UserHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/user/profile", "Profile")
	b.Handle(iris.MethodPost, "/user", "Create", middlewares.RecordSystemLog("Create", "", "创建用户成功"))
	b.Handle(iris.MethodGet, "/user/{id:string}", "Get")
	b.Handle(iris.MethodGet, "/users", "List")
	b.Handle(iris.MethodPut, "/user/{id:string}", "Update", middlewares.RecordSystemLog("Update", "id", "更新用户成功"))
	b.Handle(iris.MethodDelete, "/user/{id:string}", "Delete", middlewares.RecordSystemLog("Delete", "id", "删除用户成功"))
	b.Handle(iris.MethodDelete, "/user/multi", "MultiDelete", middlewares.RecordSystemLog("Delete", "ids", "批量删除用户成功"))
	b.Handle(iris.MethodPut, "/user/{id:string}/roles", "UpdateRoles", middlewares.RecordSystemLog("Update", "id", "更新用户角色成功"))
	b.Handle(iris.MethodPut, "/user/{id:string}/jsc_js", "UpdateJSCAndJS", middlewares.RecordSystemLog("Update", "id", "更新用户js分类成功"))
	b.Handle(iris.MethodGet, "/user/{id:string}/roles", "GetRolesByUserID")
	b.Handle(iris.MethodGet, "/user/{id:string}/jsc_js", "GetJscAndJsByUserID")
	b.Handle(iris.MethodGet, "/user/menus", "GetUserMenus")
	b.Handle(iris.MethodPatch, "/user/{id:string}", "StatusChange", middlewares.RecordSystemLog("Update", "id", "更新用户成功"))
}
