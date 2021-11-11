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
	id, err := ctx.Params().GetUint(constant.ID)
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
// @Param keywords query string false "keywords" "搜索关键词过滤"
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
	resp, ex := u.Svc.List(params)
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
	id, err := ctx.Params().GetUint(constant.ID)
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
// @Summary 删除用户(有其他依赖关系，暂不可调)
// @Description 删除用户
// @Tags 权限管理 - 管理员
// @Param id path string true "用户id"
// @Param parameters body vo.UserUpdateReq true "UserUpdateReq"
// @Success 200 "删除用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/permission/user/{id} [delete]

// BeforeActivation 初始化路由
func (u *UserHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/user/profile", "Profile")
	b.Handle(iris.MethodPost, "/user", "Create")
	b.Handle(iris.MethodGet, "/user/{id:string}", "Get")
	b.Handle(iris.MethodGet, "/users", "List")
	b.Handle(iris.MethodPut, "/user/{id:string}", "Update")
}
