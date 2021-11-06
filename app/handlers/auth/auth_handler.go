package auth

import (
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/exception"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginHandler struct {
	Svc service.LoginService
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		Svc: service.GetLoginService(),
	}
}

// Create godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags 登录
// @Param parameters body vo.LoginReq true "LoginReq"
// @Success 200 {object} vo.LoginResponse "响应成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Router /auth/login [post]
func (lh *LoginHandler) Login(ctx iris.Context) mvc.Result {
	user := &vo.LoginReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	res, ex := lh.Svc.Login(user.UserName, user.Password)
	if ex != nil {
		return response.Error(ex)
	}
	ctx.Values().Set("token", res.AccessToken)
	return response.JSON(res)
}

// Create godoc
// @Summary 用户登出
// @Description 用户登出
// @Tags 登录
// @Success 200  "响应成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Router /auth/logout [get]
func (lh *LoginHandler) Logout(ctx iris.Context) mvc.Result {
	return response.OK()
}

// BeforeActivation 初始化路由
func (u *LoginHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/login", "Login")
	b.Handle(iris.MethodGet, "/logout", "Logout")
}
