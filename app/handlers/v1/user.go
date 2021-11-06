package v1

import (
	"js_statistics/app/handlers"
	"js_statistics/app/response"
	"js_statistics/app/service"
	"js_statistics/app/vo"
	"js_statistics/exception"

	"github.com/iris-contrib/middleware/jwt"
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
// @Summary 创建用户
// @Description 创建用户
// @Tags 权限管理 - 管理员
// @Param parameters body vo.UserReq true "UserReq"
// @Success 200  "创建用户成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Router /api/v1/permission/user [post]
func (u *UserHandler) Create(ctx iris.Context) mvc.Result {
	user := &vo.UserReq{}
	if err := ctx.ReadJSON(user); err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestBody, err))
	}
	ex := u.Svc.Create(user)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// Create godoc
// @Summary 登录验证
// @Description 登录验证
// @Tags 登录验证
// @Router /api/v1/permission/print [get]
func (lh *UserHandler) Print(ctx iris.Context) {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims
	ctx.JSON(jwtInfo)
}

// BeforeActivation 初始化路由
func (u *UserHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/user", "Create")
	b.Handle(iris.MethodGet, "/print", "Print")
}
