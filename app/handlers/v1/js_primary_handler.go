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

type JspHandler struct {
	handlers.BaseHandler
	Svc service.JspService
}

func NewJsPrimaryHandler() *JspHandler {
	return &JspHandler{
		Svc: service.GetJspService(),
	}
}

// Create godoc
// @Summary 创建js主分类
// @Description 创建js主分类管理
// @Tags 应用管理 - js主分类
// @Param parameters body vo.JsPrimaryReq true "JsPrimaryReq"
// @Success 200  "创建js主分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary [post]
func (ch *JspHandler) Create(ctx iris.Context) mvc.Result {
	cdn := &vo.JsPrimaryReq{}
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
// @Summary 查询js主分类列表
// @Description 查询js主分类列表
// @Tags 应用管理 - js主分类
// @Param page query int false "请求页"
// @Param page_size query int false "页大小"
// @Param keywords query string false "keywords" "搜索关键词过滤"
// @Success 200 {object} vo.DataPagination{data=[]vo.JsPrimaryResp} "查询js主分类列表成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primaries [get]
func (ch *JspHandler) List(ctx iris.Context) mvc.Result {
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
// @Summary 查询js主分类
// @Description 查询js主分类信息
// @Tags 应用管理 - js主分类
// @Param id path string true "js主分类id"
// @Success 200 {object} vo.JsPrimaryResp "查询域名成功"
// @Failure 400 {object} vo.Error  "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/{id} [get]
func (ch *JspHandler) Get(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetUint(constant.ID)
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
// @Summary 修改js主分类
// @Description 修改js主分类信息
// @Tags 应用管理 - js主分类
// @Param id path string true "js主分类id"
// @Param parameters body vo.JsPrimaryUpdateReq true "JsPrimaryUpdateReq"
// @Success 200 "修改js主分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/{id} [put]
func (ch *JspHandler) Update(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetUint(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	param := &vo.JsPrimaryUpdateReq{}
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
// @Summary 删除js主分类
// @Description 删除js主分类信息
// @Tags 应用管理 - js主分类
// @Param id path string true "js主分类id"
// @Success 200 "删除js主分类成功"
// @Failure 400 {object} vo.Error "请求参数错误"
// @Failure 401 {object} vo.Error "当前用户登录令牌失效"
// @Failure 403 {object} vo.Error "当前操作无权限"
// @Failure 500 {object} vo.Error "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/v1/application/js_primary/{id} [delete]
func (ch *JspHandler) Delete(ctx iris.Context) mvc.Result {
	id, err := ctx.Params().GetUint(constant.ID)
	if err != nil {
		return response.Error(exception.Wrap(response.ExceptionInvalidRequestParameters, err))
	}
	ex := ch.Svc.Delete(id)
	if ex != nil {
		return response.Error(ex)
	}
	return response.OK()
}

// BeforeActivation 初始化路由
func (ch *JspHandler) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/js_primary", "Create")
	b.Handle(iris.MethodGet, "/js_primaries", "List")
	b.Handle(iris.MethodGet, "/js_primary/{id:string}", "Get")
	b.Handle(iris.MethodPut, "/js_primary/{id:string}", "Update")
	b.Handle(iris.MethodDelete, "/js_primary/{id:string}", "Delete")
}
